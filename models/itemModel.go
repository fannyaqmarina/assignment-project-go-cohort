package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string
	Description string
	OrderID     uint
	Quantity    int
}
