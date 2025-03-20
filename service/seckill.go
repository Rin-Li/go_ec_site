package service

import (
	"context"
	"gin-mall-tmp/cache"
	"gin-mall-tmp/dao"
	"gin-mall-tmp/model"
	"gin-mall-tmp/pkg/e"
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/serializer"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type SeckillService struct {
}

func (service *SeckillService) ShowProduct(ctx context.Context) serializer.Response {
	var products []model.SeckillProduct
	var err error
	code := e.Success
	productDao := dao.NewSeckillDao(ctx)
	products, err = productDao.GetSeckillProducts()
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildSeckillProducts(products), uint(len(products)))

}

func (service *SeckillService) Order(ctx context.Context, uId uint, pid string) serializer.Response {
	id, _ := strconv.Atoi(pid)
	code := e.Success

	seckillproductDao := dao.NewSeckillDao(ctx)
	product, err := cache.GetSeckillProduct(uint(id))

	if err != nil {
		if err == redis.Nil {
			product, err = seckillproductDao.GetSeckillProductById(uint(id))
			if err != nil {
				code = e.ErrorNotSeckillProduct
				return serializer.Response{
					Status: code,
					Data:   e.GetMsg(code),
					Error:  err.Error(),
				}
			}
			go func() {
				cache.SetSeckillProduct(uint(id), product)
			}()
		} else {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

	}

	if product.Stock <= 0 {
		code = e.ErrorOutOfStock
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}

	}

	tx := seckillproductDao.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = tx.Model(&model.SeckillProduct{}).Where("id=?", id).Update("stock", product.Stock-1).Error
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	order := model.SeckillOrder{
		UserID:      uId,
		ProductID:   uint(id),
		OrderStatus: 0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		code = e.ErrorOrderSeckill
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	tx.Commit()

	go func() {
		product.Stock -= 1
		cache.SetSeckillProduct(uint(id), product)
	}()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildSeckillOrder(order),
	}
}
