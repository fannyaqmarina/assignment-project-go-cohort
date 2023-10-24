package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()
	r.POST("/order", controllers.OrderCreate)
	r.GET("/order", controllers.OrderGet)
	r.GET("/order/:orderId", controllers.OrderGetOne)
	r.PUT("/order/:orderId", controllers.OrderUpdate)
	r.DELETE("/order/:orderId", controllers.OrderDelete)
	r.Run()
}
