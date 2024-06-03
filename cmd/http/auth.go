package http

import (
	"awesome-auth/internal/core/http/middlewares"
	"awesome-auth/internal/core/services/auth"
	"awesome-auth/internal/repositories"

	"github.com/gin-gonic/gin"
)

func (s *Server) DefineAuthRoutes(router *gin.RouterGroup) {

	repo := repositories.NewUserRepo(&s.DB)
	tokenRepo := repositories.NewTokenRepo(&s.DB)
	service := auth.NewAuthService(repo, tokenRepo)

	authRouter := router.Group("auth")

	middlewares.CommonMiddlewares(router)

	{
		authRouter.POST("login", service.Login)
		authRouter.POST("logout", service.Logout)
		authRouter.POST("register", service.Register)
		authRouter.GET("verify", service.Verify).Use(middlewares.CheckTokenExists())
		authRouter.GET("me", service.GetMe).Use(middlewares.CheckTokenExists())
	}
}
