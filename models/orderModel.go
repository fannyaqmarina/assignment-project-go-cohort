package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Item         []Item
	CustomerName string
	OrderedAt    time.Time
}
