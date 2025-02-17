package v1

import (
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ListSeckillProducts(c *gin.Context){
	var listSeckillProducts service.SeckillService
	if err := c.ShouldBind(&listSeckillProducts);err == nil{
		res := listSeckillProducts.ShowProduct(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
	}
}

func SeckillOrder(c *gin.Context){
	var seckillOrder service.SeckillService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&seckillOrder);err == nil{
		res := seckillOrder.Order(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
	}
}


