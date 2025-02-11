package v1

import (
	"gin-mall-tmp/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context){
	var listCategory service.CategoryService
	if err := c.ShouldBind(&listCategory);err == nil{
		res := listCategory.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
	}
}