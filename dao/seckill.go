package dao

import (
	"context"
	"gin-mall-tmp/model"

	"gorm.io/gorm"
)

type SeckillDao struct {
	*gorm.DB 
}

func NewSeckillDao(ctx context.Context) *SeckillDao {
	return &SeckillDao{NewDBClient(ctx)}
}

func NewSeckillDaoByDB(db *gorm.DB) *SeckillDao {
	return &SeckillDao{db}
}

func (dao *SeckillDao) GetSeckillProducts() ([]model.SeckillProduct, error){
	var products []model.SeckillProduct
	err := dao.DB.Find(&products).Error
	return products, err
}

func (dao *SeckillDao) GetSeckillProductById(id uint) (*model.SeckillProduct, error){
	var product model.SeckillProduct
	err := dao.DB.Where("id=?", id).First(&product).Error
	return &product, err
}