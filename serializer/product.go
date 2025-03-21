package serializer

import (
	"gin-mall-tmp/cache"
	"gin-mall-tmp/conf"
	"gin-mall-tmp/model"
)

type Product struct{
	Id uint `json:"id"`
	Name string `json:"name"`
	CategoryId uint `json:"category_id"`
	Title string `json:"title"`
	Info string `json:"info"`
	ImgPath string `json:"img_path"`
	Price string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View uint64 `json:"view"`
	CreateAt int64 `json:"create_at"`
	Num int `json:"num"`
	OnSale bool `json:"on_sale"`
	BossId uint `json:"boss_id"`
	BossName string `json:"boss_name"`
	BossAvatar string `json:"boss_avatar"`
}

func BuildProduct(item model.Product) Product {
	viewCount := cache.View(item.ID)

	return Product{
		Id: item.ID,
		Name: item.Name,
		CategoryId: item.CategoryId,
		Title: item.Title,
		Info: item.Info,
		ImgPath: conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath, 
		Price: item.Price,
		DiscountPrice: item.DiscountPrice,
		View: viewCount,
		CreateAt: item.CreatedAt.Unix(),
		Num: item.Num,
		OnSale: item.Onsale,
		BossId: item.BossId, 
		BossName: item.BossName,
		BossAvatar: item.BossAvatar,
	}
}

func BuildProducts(items []model.Product) (products []Product){
	for _, item := range items{
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}

