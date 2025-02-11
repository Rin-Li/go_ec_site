package v1

import (
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

//Create product
func CreateFavorites(c *gin.Context){
	claim,_ := util.ParseToken(c.GetHeader("Authorization"))
	createFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&createFavoriteService);err == nil{
		res := createFavoriteService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func ListFavorites(c *gin.Context){
	claim,_ := util.ParseToken(c.GetHeader("Authorization"))
	listFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&listFavoriteService);err == nil{
		res := listFavoriteService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func DeleteFavorites(c *gin.Context){
	claim,_ := util.ParseToken(c.GetHeader("Authorization"))
	deleteFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&deleteFavoriteService);err == nil{
		res := deleteFavoriteService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

