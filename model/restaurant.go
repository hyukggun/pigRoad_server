package model

import (
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Rating    float32 `json:"rating"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Menus     []Menu  `json:"menus"`
}
