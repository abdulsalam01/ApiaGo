package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Id      uint   `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Address string `gorm:"not null" json:"address"`
	User    `gorm:"foreignKey: id"`
}
