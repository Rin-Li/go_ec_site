package model

import "gorm.io/gorm"

type Order struct{
	gorm.Model
	UserId uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId uint `gorm:"not null"`
	AddressId uint `gorm:"not null"`
	Num int
	OrderNum uint64
	Type uint //1 not paid 2 has paid
	Money float64
}