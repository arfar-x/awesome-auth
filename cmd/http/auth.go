package http

import (
	"awesome-auth/internal/core/service"
	"awesome-auth/internal/mysql"

	"github.com/gin-gonic/gin"
)

func (s *Server) DefineAuthRoutes(router *gin.RouterGroup) {

	repo := mysql.NewUserRepo(&s.DB)
	srv := service.NewAuthService(repo)

	auth := router.Group("auth")

	auth.POST("login", srv.Login)
	auth.POST("logout", srv.Logout)
	auth.POST("register", srv.Register)
	auth.POST("verify", srv.Verify)
	auth.GET("me", srv.GetMe)
}
