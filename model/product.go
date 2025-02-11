package model

import (
	"gin-mall-tmp/cache"
	"strconv"

	"gorm.io/gorm"
)

type Product struct{
	gorm.Model
	Name string
	CategoryId uint
	Title string
	Info string
	ImgPath string
	Price string
	DiscountPrice string
	Onsale bool `gorm:"deafult:false"`
	Num int
	BossId uint
	BossName string
	BossAvatar string
}

func (product *Product) View() uint64{
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product)AddView(){
	//Increase product click
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa((int(product.ID))))
	
}