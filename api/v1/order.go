package v1

import (
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create product
func CreateOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createOrderervice := service.OrderService{}
	if err := c.ShouldBind(&createOrderervice); err == nil {
		res := createOrderervice.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func ListOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	listOrderervice := service.OrderService{}
	if err := c.ShouldBind(&listOrderervice); err == nil {
		res := listOrderervice.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func DeleteOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	deleteOrderervice := service.OrderService{}
	if err := c.ShouldBind(&deleteOrderervice); err == nil {
		res := deleteOrderervice.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func ShowOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showOrderervice := service.OrderService{}
	if err := c.ShouldBind(&showOrderervice); err == nil {
		res := showOrderervice.Show(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

