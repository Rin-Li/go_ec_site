package middleware

import (
	"gin-mall-tmp/pkg/e"
	"gin-mall-tmp/pkg/util"
	"time"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context){
		var code int
		code = 200
		token := c.GetHeader("Authorization")

		if token ==""{
			code = 404
		} else{
			claims, err := util.ParseToken(token)
			if err != nil{
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt{
				code = e.ErrorAuthCheckTokenTimeOut
			}
		}
		if code != e.Success{
			c.JSON(200, gin.H{
				"status":code,
				"msg":e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}