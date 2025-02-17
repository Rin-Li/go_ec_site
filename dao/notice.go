package dao

import (
	"context"
	"gin-mall-tmp/model"

	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB 
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticerDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}
//Get notice by Id
func (dao *NoticeDao) GetNoticeById(id uint)(notice *model.Notice, err error){
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
} 
