package serializer

import (
	"context"
	"gin-mall-tmp/conf"
	"gin-mall-tmp/dao"
	"gin-mall-tmp/model"
)

type Order struct{
	Id uint `json:"id"`
	OrderNum uint64 `json:"order_num"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
	UserId uint `json:"user_id"`
	ProductId uint `json:"product_id"`
	BossId uint `json:"boss_id"`
	Num int `json:"num"`
	AddressName string `json:"address_name"`
	AddressPhone string `json:"address_phone"`
	Address string `json:"address"`
	Type uint `json:"type"`
	ProductName string `json:"product_name"`
	ImgPath string `json:"img_path"`
	Money float64 `json:"discount_price"`
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address) Order {

    return Order{
		Id: order.ID,
		OrderNum: order.OrderNum,
		CreateAt: order.CreatedAt.Unix(),
		UpdateAt: order.UpdatedAt.Unix(),
		UserId: order.UserId,
		ProductId: order.ProductId,
		BossId: order.BossId,
		Num: order.Num,
		AddressName: address.Name,
		AddressPhone: address.Phone,
		Address: address.Address,
		Type: order.Type,
		ProductName: product.Name,
		ImgPath: conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Money: order.Money,
    }
}

func BuildOrders(ctx context.Context, items []*model.Order) []Order {
	productDao := dao.NewProductDao(ctx)
	addressDao := dao.NewAddressDao(ctx)

	var orders []Order
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		address, err := addressDao.GetAddressByAid(item.AddressId)
		if err != nil {
			continue
		}
		order := BuildOrder(item, product, address) 
		orders = append(orders, order)
	}
	return orders
}

