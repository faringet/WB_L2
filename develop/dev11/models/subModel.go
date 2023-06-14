package models

import "gorm.io/gorm"

type Sub struct {
	gorm.Model
	City        string
	Temperature float64
}
