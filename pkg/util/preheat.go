package util

import (
	"context"
	"gin-mall-tmp/cache"
	"gin-mall-tmp/dao"
)


func PreheatSeckillProducts(){
	ctx := context.Background()
	serkillDao := dao.NewSeckillDao(ctx)

	products, err := serkillDao.GetSeckillProducts()
	if err != nil{
		LogrusObj.Infoln(err)
		return
	}

	successCount := 0
	failedCount := 0

	for _, product := range products{
		err := cache.SetSeckillProduct(product.ID, &product)
		if err != nil{
			LogrusObj.Printf("[Preheat] prodct %d failed: %v\n", product.ID, err)
			failedCount++
		} else{
			successCount++
		}
	}

	LogrusObj.Printf("[Preheat] success: %d, failed: %d\n", successCount, failedCount)

}