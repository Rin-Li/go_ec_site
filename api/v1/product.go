package v1

import (
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Create product
func CreateProduct(c *gin.Context){
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim,_ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := c.ShouldBind(&createProductService);err == nil{
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func ListProduct(c *gin.Context){
	listProductService := service.ProductService{}
	if err := c.ShouldBind(&listProductService);err == nil{
		res := listProductService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}

func ShowProduct(c *gin.Context){
	showProductService := service.ProductService{}
	if err := c.ShouldBind(&showProductService);err == nil{
		res := showProductService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
} 



func SearchProduct(c *gin.Context){
	searchProductService := service.ProductService{}
	if err := c.ShouldBind(&searchProductService);err == nil{
		res := searchProductService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResPonse(err))
	}
}
