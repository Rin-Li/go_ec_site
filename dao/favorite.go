package dao

import (
	"context"
	"gin-mall-tmp/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}


func (dao *FavoriteDao) ListFavorite(uId uint) (favorites []*model.Favorite, err error) {
	err = dao.DB.Where("user_id = ?", uId).Find(&favorites).Error
	return
}


func (dao *FavoriteDao) FavoriteExistOrNot(pId, uId uint) (bool, error) {
	var count int64
	err := dao.DB.Where("product_id = ? and user_id = ?", pId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (dao *FavoriteDao) CreateFavorite(in *model.Favorite) error{
	return dao.DB.Model(&model.Favorite{}).Create(&in).Error
}

func (dao *FavoriteDao) DeleteFavorite(uId uint, fId uint) error{
	return dao.DB.Model(&model.Favorite{}).Where("user_id = ? and id = ?", uId, fId).Delete(&model.Favorite{}).Error
}
