package v1

import (
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create product
func CreateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createAddresservice := service.AddressService{}
	if err := c.ShouldBind(&createAddresservice); err == nil {
		res := createAddresservice.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func ListAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	listAddresservice := service.AddressService{}
	if err := c.ShouldBind(&listAddresservice); err == nil {
		res := listAddresservice.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func DeleteAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	deleteAddresservice := service.AddressService{}
	if err := c.ShouldBind(&deleteAddresservice); err == nil {
		res := deleteAddresservice.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func ShowAddress(c *gin.Context) {
	showAddresservice := service.AddressService{}
	if err := c.ShouldBind(&showAddresservice); err == nil {
		res := showAddresservice.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func UpdateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateAddresservice := service.AddressService{}
	if err := c.ShouldBind(&updateAddresservice); err == nil {
		res := updateAddresservice.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

