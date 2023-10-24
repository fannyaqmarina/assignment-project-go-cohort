package controllers

import (
	"fmt"
	"go-crud/initializers"
	"go-crud/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func OrderCreate(c *gin.Context) {

	var bodyOrder struct {
		Item         []models.Item
		CustomerName string
		OrderedAt    time.Time
	}

	c.Bind(&bodyOrder)

	order := models.Order{CustomerName: bodyOrder.CustomerName, Item: bodyOrder.Item, OrderedAt: bodyOrder.OrderedAt}

	fmt.Println("ORDER : ", order)

	result := initializers.DB.Create(&order)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": order,
	})
}

func OrderGet(c *gin.Context) {
	var orders []models.Order
	if err := initializers.DB.Preload("Item").Find(&orders).Error; err != nil {
		fmt.Errorf(err.Error())
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"order": orders,
	})
}

func OrderGetOne(c *gin.Context) {
	orderId, _ := c.Params.Get("orderId")
	var order models.Order
	if err := initializers.DB.Preload("Item").First(&order, orderId).Error; err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.Errors.JSON()
			c.Status(404)
			return
		}
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"order": order,
	})
}

func OrderUpdate(c *gin.Context) {
	orderId, _ := c.Params.Get("orderId")
	var order models.Order
	if err := initializers.DB.Preload("Item").Where("id = ?", orderId).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}
	var bodyOrder struct {
		OrderedAt    time.Time     `json:"orderedAt"`
		CustomerName string        `json:"customerName"`
		Items        []models.Item `json:"items"`
	}
	if err := c.ShouldBindJSON(&bodyOrder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Model(&order).Updates(models.Order{CustomerName: bodyOrder.CustomerName, OrderedAt: bodyOrder.OrderedAt})

	for i := 0; i < len(order.Item); i++ {
		if err := initializers.DB.Model(&models.Item{}).Where("id = ?", order.Item[i].ID).Updates(models.Item{Name: bodyOrder.Items[i].Name, Description: bodyOrder.Items[i].Description, Quantity: bodyOrder.Items[i].Quantity}).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update item"})
			return
		}
	}
	c.JSON(200, gin.H{
		"order": order,
	})
}

func OrderDelete(c *gin.Context) {
	orderId, _ := c.Params.Get("orderId")
	var order models.Order
	if err := initializers.DB.Preload("Item").Where("id = ?", orderId).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	for i := 0; i < len(order.Item); i++ {

		if err := initializers.DB.Delete(&models.Item{}, order.Item[i].ID).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete item"})
			return
		}
	}
	if err := initializers.DB.Delete(&models.Order{}, orderId).Error; err != nil {
		fmt.Println("Error : ", err)
		c.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}
	c.JSON(200, gin.H{
		"orderId": orderId,
	})
}
