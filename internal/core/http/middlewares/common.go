package middlewares

import (
	"awesome-auth/pkg/response"
	"github.com/gin-gonic/gin"
)

func CommonMiddlewares(router *gin.RouterGroup) {
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
}

func CheckTokenExists() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := ctx.Request.Header["Authorization"]; !ok {
			ctx.AbortWithStatusJSON(response.Unauthorized("Unauthorized middleware.", nil))
			return
		}

		ctx.Next()
	}
}
