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
	r.POST("/orders", controllers.OrderCreate)
	r.GET("/orders", controllers.OrderGet)
	r.GET("/orders/:orderId", controllers.OrderGetOne)
	r.PUT("/orders/:orderId", controllers.OrderUpdate)
	r.DELETE("/orders/:orderId", controllers.OrderDelete)
	r.Run()
}
