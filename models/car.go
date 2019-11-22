package models

import (
	"github.com/jinzhu/gorm"
)

type Car struct {
	gorm.Model
	Owner          string `json:"owner"`
	Color          string `json:"color"`
	NumberOfWheels int    `json:"number_of_wheels"`
}
