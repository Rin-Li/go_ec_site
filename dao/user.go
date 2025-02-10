package dao

import (
	"context"
	"gin-mall-tmp/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB //表示与数据库的连接
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 根据username判断用户是否存在
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).Find(&user).Count(&count).Error
	if count == 0{
		return nil, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Create(user).Error
}

//Get user byd id
func (dao *UserDao) GetUserById(id uint)(user *model.User, err error){
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
} 

//UpdateUserByid
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) error{
	return dao.DB.Model(&model.User{}).Where("id=?", uId).Updates(&user).Error
}