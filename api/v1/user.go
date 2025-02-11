package v1

import (
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context){
	var UserRegister service.UserService
	if err := c.ShouldBind(&UserRegister);err == nil{
		res := UserRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UserLogin(c *gin.Context){
	var UserLogin service.UserService
	if err := c.ShouldBind(&UserLogin);err == nil{
		res := UserLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
		util.LogrusObj.Infoln(err)
	}
}

func Userupdate(c *gin.Context){
	var UserUpdate service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization")) //First valida the token
	if err := c.ShouldBind(&UserUpdate);err == nil{
		res := UserUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UploadAvatar(c *gin.Context){
	file, fileHeader,_ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization")) //First valida the token
	if err := c.ShouldBind(&uploadAvatar);err == nil{
		res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
		util.LogrusObj.Infoln(err)
	}
	
}

func SendEmail(c *gin.Context){
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(c.GetHeader("Authorization")) //First valida the token
	if err := c.ShouldBind(&sendEmail);err == nil{
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
		util.LogrusObj.Infoln(err)
	}

}

func ValidEmail(c *gin.Context){
	var ValidEmail service.ValidEmailService

	if err := c.ShouldBind(&ValidEmail);err == nil{
		res := ValidEmail.Valid(c.Request.Context(), c.GetHeader("Authorization"))
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
		util.LogrusObj.Infoln(err)
	}

}

func ShowMoeny(c *gin.Context){
	var showMoeny service.ShowMoenyService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoeny);err == nil{
		res := showMoeny.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest,ErrorResPonse(err))
		util.LogrusObj.Infoln(err)
	}

}