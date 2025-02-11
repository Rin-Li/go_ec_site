package service

import (
	"context"
	"errors"
	"fmt"
	"gin-mall-tmp/dao"
	"gin-mall-tmp/model"
	"gin-mall-tmp/pkg/e"
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/serializer"
	"strconv"
)

type OrderPay struct{
	OrderId uint `json:"order_id" form:"order_id"`
	Money float64 `json:"money" form:"money"`
	OrderNo string `json:"order_no" form:"order_no"`
	ProductId uint `json:"product_id" form:"product_id"`
	PayTime string `json:"pay_time" form:"pay_time"`
	Sign string `json:"sign" form:"sign"`
	BossId uint `json:"boss_id" form:"boss_id"`
	BossName string `json:"boss_name" form:"boss_name"`
	Num int `json:"num" form:"num"`
	Key string `json:"key" form:"key"` //Pay amount
}

func (service *OrderPay) PayDown(ctx context.Context, uId uint) serializer.Response{
	util.Encrypt.SetKey(service.Key)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)

	tx := orderDao.Begin()
	order, err := orderDao.GetOrderByid(service.OrderId, uId)
	if err != nil{
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)

	if err != nil{
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	//Get moeny, then substract pay amout, save again, User pay
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	
	if moneyFloat - money < 0.0{
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: errors.New("Insufficient balance").Error(),
		}
	}
	finMoney := fmt.Sprintf("%f", moneyFloat - money)
	user.Money = util.Encrypt.AesEncoding(finMoney)

	userDao = dao.NewUserDaoByDB(userDao.DB)
	err = userDao.UpdateUserById(uId, user)

	if err != nil{
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	//Solder get money
	var boss *model.User
	boss, err = userDao.GetUserById(service.BossId)
	if err != nil{
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMondy := fmt.Sprintf("%f", moneyFloat + money)
	boss.Money = util.Encrypt.AesEncoding(finMondy)

	err = userDao.UpdateUserById(service.BossId, boss)
	if err != nil{
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	//Product num -1
	var product *model.Product
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(service.ProductId)
	product.Num = product.Num - 1

	if err != nil{
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	err = productDao.UpdateProduct(service.ProductId, product)
	if err != nil{
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	//Order status change
	order.Type = 2
	err = orderDao.UpdateOrderByOrderId(service.OrderId, order)
	if err != nil{
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	//Own product + 1
	productUser := model.Product{
		Name: product.Name,
		CategoryId: product.CategoryId,
		Title: product.Title,
		Info: product.Info,
		ImgPath: product.ImgPath,
		Price: product.Price,
		DiscountPrice: product.DiscountPrice,
		Onsale: false,
		Num: 1,
		BossId: uId,
		BossName: user.UserName,
		BossAvatar: user.Avater,
	}
	err = productDao.CreateProduct(&productUser)
	
	if err != nil{
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	tx.Commit()
	return serializer.Response{ 
		Status: code,
		Msg: e.GetMsg(code),
	}
}