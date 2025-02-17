package model

import (
	"time"
)

type SeckillProduct struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"index" json:"product_id"` 
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	StartTime time.Time `json:"start_time"` 
	EndTime   time.Time `json:"end_time"`   
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SeckillOrder struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"index" json:"user_id"`
	ProductID   uint      `gorm:"index" json:"product_id"`
	OrderStatus uint   `json:"order_status"` // 0: unpaid, 1: paid
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

