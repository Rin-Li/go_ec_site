package v1

import (
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

func OrderPay(c *gin.Context) {
	
	orderPay := service.OrderPay{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&orderPay); err == nil {
		res := orderPay.PayDown(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}