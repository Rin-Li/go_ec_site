package dao

import (
	"context"
	"gin-mall-tmp/model"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error{
	return dao.DB.Model(&model.Order{}).Create(&in).Error
}

func (dao *OrderDao) GetOrderByid(oId uint, uId uint) (order *model.Order, err error){
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", oId, uId).First(&order).Error
	return
}

func (dao *OrderDao) ListOrderByUserId(uId uint) (orderes []*model.Order, err error){
	err = dao.DB.Where("user_id = ?", uId).Find(&orderes).Error
	return
}

func (dao *OrderDao) UpdateOrderByOrderId(oId uint, order *model.Order) error{
	return dao.DB.Model(&model.Order{}).Where("id = ?", oId).Updates(&order).Error
}

func (dao *OrderDao) DeleteOrderByOrderId(aId uint, uId uint) error{
	return dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ?", aId, uId).Delete(&model.Order{}).Error
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, err error){
	err = dao.DB.Model(&model.Order{}).Where(condition).
	Offset((page.PageNum - 1)*(page.PageSize)).Limit(page.PageSize).
	Find(&orders).Error
	return
}