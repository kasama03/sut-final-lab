package entity

import (
	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	Name       string
	Age        string `valid: "IsInt~Age is not integer"`
	CustomerID string
}
