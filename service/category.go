package service

import (
	"context"
	"gin-mall-tmp/dao"
	"gin-mall-tmp/serializer"
	"gin-mall-tmp/pkg/e"
	"gin-mall-tmp/pkg/util"
)

type CategoryService struct {

}

func (service *CategoryService) List(ctx context.Context) serializer.Response{
	categoryDao := dao.NewCategoryDao(ctx)
	
	code := e.Success
	categorys, err := categoryDao.ListCategory()
	if err != nil{
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategories(categorys), uint(len(categorys)))
}