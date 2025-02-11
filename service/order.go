package service

import (
	"context"
	"fmt"
	"gin-mall-tmp/dao"
	"gin-mall-tmp/model"
	"gin-mall-tmp/pkg/e"
	"gin-mall-tmp/serializer"
	"math/rand"
	"strconv"
	"time"
	
)

type OrderService struct {
	ProductId uint `json:"product_id" form:"product_id"`
	Num int `json:"num" form:"num"`
	AddressId uint `json:"address_id" form:"address_id"`
	Money float64 `json:"money" form:"money"`
	BossId uint `json:"boss_id" form:"boss_id"`
	UserId uint `json:"user_id" form:"user_id"`
	OrderNum int `json:"order_num" form:"order_num"`
	Type int `json:"type" form:"type"` //1 not paid 2 has paid
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	var order model.Order
	code := e.Success
	OrderDao := dao.NewOrderDao(ctx)
	order = model.Order{
		UserId: uId,
		ProductId: service.ProductId,
		BossId: service.BossId,
		Num: service.Num,
		Money: service.Money,
		Type: 1, //Unpaid
	}
	//valid address exit
	adddressDao := dao.NewAddressDao(ctx)
	address, err := adddressDao.GetAddressByAid(service.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	
	//Generate random number and unique order number
	order.AddressId = address.ID
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000)) //order_number
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	order.OrderNum, _ = strconv.ParseUint(number, 10, 64)

	err = OrderDao.CreateOrder(&order)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
}

func (service *OrderService) Show(ctx context.Context, uId uint, oId string) serializer.Response {
	orderId, _ := strconv.Atoi(oId)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderByid(uint(orderId), uId)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(order.AddressId)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data: serializer.BuildOrder(order, product, address),
	}
}

func (service *OrderService) List(ctx context.Context, uId uint) serializer.Response{
	code := e.Success

	if service.PageSize == 0{
		service.PageSize = 15
	}

	OrderDao := dao.NewOrderDao(ctx)

	condition := make(map[string]interface{})
	if service.Type != 0{
		condition["type"] = service.Type // All orders
	}

	orderList, err := OrderDao.ListOrderByCondition(condition, service.BasePage)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data: serializer.BuildOrders(ctx, orderList),
	}
}


func (service *OrderService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	OrderId, _ := strconv.Atoi(aId)
	code := e.Success
	OrderDao := dao.NewOrderDao(ctx)
	err := OrderDao.DeleteOrderByOrderId(uint(OrderId), uId)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
}