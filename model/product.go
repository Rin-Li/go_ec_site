package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	Onsale        bool `gorm:"deafult:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}


