package serializer

import (
	"context"
	"gin-mall-tmp/conf"
	"gin-mall-tmp/dao"
	"gin-mall-tmp/model"
)

type Favorite struct {
	UserId uint `json:"user_id"`
	ProductId uint `json:"product_id"`
	CreateAt int64 `json:"create_at"`
	Name string `json:"name"`
	CategoryId uint `json:"category_id"`
	Title string `json:"title"`
	Info string `json:"info"`
	ImgPath string `json:"img_path"`
	Price string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossId uint `json:"boss_id"`
	Num int `json:"num"`
	OnSale bool `json:"on_sale"`
}

func BuildFavorite(favorite *model.Favorite, product *model.Product, boss *model.User) Favorite {
    return Favorite{
		UserId: favorite.UserId,
		ProductId: favorite.ProductId,
		CreateAt: favorite.CreatedAt.Unix(),
		Name: product.Name,
		CategoryId: product.CategoryId,
		Title: product.Title,
		Info: product.Info,
		ImgPath: conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price: product.Price,
		DiscountPrice: product.DiscountPrice,
		BossId: boss.ID,
		Num: product.Num,
		OnSale: product.Onsale,
    }
}

func BuildFavorites(ctx context.Context, items []*model.Favorite) []Favorite {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)

	var favorites []Favorite
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.UserId)
		if err != nil {
			continue
		}
		favorite := BuildFavorite(item, product, boss) // item 本身是指针
		favorites = append(favorites, favorite)
	}
	return favorites
}

