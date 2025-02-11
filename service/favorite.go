package service

import (
	"context"
	"gin-mall-tmp/dao"
	"gin-mall-tmp/model"
	"gin-mall-tmp/pkg/e"
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/serializer"
	"strconv"
)

type FavoriteService struct{
	ProductId uint `json:"product_id" form:"product_id"`
	BossId uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}



func (service *FavoriteService) List(ctx context.Context, uId uint) serializer.Response{
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	favorite, err := favoriteDao.ListFavorite(uId)
	if err != nil{
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorite), uint(len(favorite)))
}

func (service *FavoriteService) Create(ctx context.Context, uId uint) serializer.Response{
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	exit, _ := favoriteDao.FavoriteExistOrNot(service.ProductId, uId)
	if exit{
		code = e.ErrorFavoriteExist
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil{
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil{
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil{
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	favorite := &model.Favorite{
		User: *user,
		UserId: uId,
		Product: *product,
		ProductId: service.ProductId,
		Boss: *boss,
		BossId: service.BossId,
	}

	favoriteDao.CreateFavorite(favorite)
	if err != nil{
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
}

func (service *FavoriteService) Delete(ctx context.Context, uId uint, fId string) serializer.Response{
	newfId, err := strconv.Atoi(fId)
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	err = favoriteDao.DeleteFavorite(uId, uint(newfId))
	if err != nil{
		util.LogrusObj.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
	
	
}