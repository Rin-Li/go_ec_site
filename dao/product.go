package dao

import (
	"context"
	"gin-mall-tmp/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB 
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}


func (dao *ProductDao) CreateProduct(product *model.Product) (err error){
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
   err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
   return
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage)(products []model.Product, err error){
	err = dao.DB.Where(condition).Offset((page.PageNum - 1)*(page.PageSize)).Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *ProductDao) SearchProduct(info string, page model.BasePage)(products []model.Product, count int64, err error){
	err = dao.DB.Model(&model.Product{}).
	Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").Count(&count).Error
	if err != nil{
		return
	}
	err = dao.DB.Model(&model.Product{}).
	Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
	Offset((page.PageNum - 1)*(page.PageSize)).Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *ProductDao) GetProductById(id uint)(product *model.Product, err error){
	err = dao.DB.Model(&model.Product{}).Where("id=?", id).First(&product).Error
	return
}

func (dao *ProductDao) UpdateProduct(id uint, product *model.Product) error{
	return dao.DB.Model(&model.Product{}).Where("id=?", id).Updates(&product).Error
}