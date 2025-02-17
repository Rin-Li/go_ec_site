package dao

import(
	"fmt"
	"gin-mall-tmp/model"
)

func Migration(){
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
	    AutoMigrate(
			&model.User{},
			&model.Address{},
			&model.Notice{},
			&model.Category{},
			&model.Carousel{},
			&model.Cart{},
			&model.Category{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.ProductImg{},
			&model.Product{},
			&model.SeckillProduct{},
			&model.SeckillOrder{},
		)
		
	if err != nil{
		fmt.Println("err", err)
	}

}