package dao

import (
	"context"
	"gin-mall-tmp/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB //表示与数据库的连接
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}
//Get notice by Id
func (dao *CategoryDao) ListCategory()(Category []model.Category, err error){
	err = dao.DB.Model(&model.Category{}).Find(&Category).Error
	return
} 