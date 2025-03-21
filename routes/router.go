package routes

import (
	api "gin-mall-tmp/api/v1"
	"gin-mall-tmp/middleware"
	"net/http"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context){
			c.JSON(200, "success")
		})
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		v1.GET("carousels", api.ListCarousel)

		//Product operation
		v1.GET("products", api.ListProduct)
		v1.GET("product/:id", api.ShowProduct)
		v1.GET("imgs/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategory)



		authed := v1.Group("/") //Need login protect
		authed.Use(middleware.JWT()) // JWT authen
		{
			//User opration
			authed.PUT("user", api.Userupdate)
			authed.POST("avater", api.UploadAvatar)
			authed.POST("user/sending_email", api.SendEmail)
			authed.POST("user/valid_email", api.ValidEmail)
			//Show amount
			authed.POST("money", api.ShowMoeny)

			//Product operation
			authed.POST("product", api.CreateProduct)
			authed.POST("products", api.SearchProduct)

			//Collection
			authed.GET("collection", api.ListFavorites)
			authed.POST("collection", api.CreateFavorites)
			authed.DELETE("collection", api.DeleteFavorites)

			//Address
			authed.POST("address", api.CreateAddress)
			authed.GET("address/:id", api.ShowAddress)
			authed.GET("address", api.ListAddress)
			authed.PUT("address/:id", api.UpdateAddress)
			authed.DELETE("address/:id", api.DeleteAddress)

			//Order
			authed.POST("order", api.CreateOrder)
			authed.GET("orders", api.ListOrder)
			authed.GET("order/:id", api.ShowOrder)
			authed.DELETE("order/:id", api.DeleteOrder)

			//Pay
			authed.POST("paydown", api.OrderPay)

			//Seckill
			authed.GET("seckill", api.ListSeckillProducts)
			authed.POST("seckill/:id", api.SeckillOrder)



		}

	}
	return r
}