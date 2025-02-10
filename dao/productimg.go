package dao

import (
	"context"
	"gin-mall-tmp/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB //表示与数据库的连接
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}


func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error{
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

func (dao *ProductImgDao) ListProductImg(id uint) (productImg []*model.ProductImg, err error){
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id=?", id).Find(&productImg).Error
	return
}