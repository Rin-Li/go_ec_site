package cache

import (
	"encoding/json"
	"errors"
	"gin-mall-tmp/model"
	"time"

	"github.com/go-redis/redis"
)


func GetSeckillProduct(productid uint) (*model.SeckillProduct, error){
	cacheKey := SeckillProductKey(productid)

	productJSON, err := RedisClient.Get(cacheKey).Result()
	if err == redis.Nil{
		return &model.SeckillProduct{}, errors.New("product not found")
	}
	if err != nil{
		return &model.SeckillProduct{}, err
	}

	var product *model.SeckillProduct
	err = json.Unmarshal([]byte(productJSON), &product)
	if err != nil{
		return &model.SeckillProduct{}, err
	}
	return product, nil

}

func SetSeckillProduct(productID uint, product *model.SeckillProduct) error {
	cacheKey := SeckillProductKey(productID)


	productJSON, err := json.Marshal(product)
	if err != nil {
		return err
	}


	err = RedisClient.Set(cacheKey, productJSON, 10*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}