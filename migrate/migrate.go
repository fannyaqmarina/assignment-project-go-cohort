package main

import (
	"go-crud/initializers"
	"go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Order{}, &models.Item{})
}
