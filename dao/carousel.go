package dao

import (
	"context"
	"gin-mall-tmp/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB //表示与数据库的连接
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}
//Get notice by Id
func (dao *CarouselDao) ListCarousel()(carousel *[]model.Carousel, err error){
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
} 
