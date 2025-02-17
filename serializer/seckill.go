package serializer

import (
	"gin-mall-tmp/model"
	"time"
)


type SeckillProduct struct {
	ID           uint   `json:"id"`
	ProductID    uint   `json:"product_id"`
	Name         string `json:"name"`
	Price        float64 `json:"price"`
	Stock        int    `json:"stock"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
}

func BuildSeckillProduct(item model.SeckillProduct) SeckillProduct{
	return SeckillProduct{
		ID: item.ID,
		ProductID: item.ProductID,
		Name: item.Name,
		Price: item.Price,
		Stock: item.Stock,
		StartTime: item.StartTime.Unix(),
		EndTime: item.EndTime.Unix(),
	}
}

func BuildSeckillProducts(items []model.SeckillProduct) (seckillProducts []SeckillProduct){
	for _, item := range items{
		seckillProduct := BuildSeckillProduct(item)
		seckillProducts = append(seckillProducts, seckillProduct)
	}
	return seckillProducts
}

type SeckillOrder struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	ProductID   uint   `json:"product_id"`
	OrderStatus uint `json:"order_status"`
	CreatedAt   int64  `json:"created_at"`
	UpdateAt    int64  `json:"updated_at"`
}

func BuildSeckillOrder(order model.SeckillOrder) SeckillOrder{
	return SeckillOrder{
		ID: order.ID,
		UserID: order.UserID,
		ProductID: order.ProductID,
		OrderStatus: order.OrderStatus,
		CreatedAt: order.CreatedAt.Unix(),
		UpdateAt: time.Now().Unix(),
	}
}

