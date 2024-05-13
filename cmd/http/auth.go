package http

import (
	"awesome-auth/internal/core/services/auth"
	"awesome-auth/internal/repositories"

	"github.com/gin-gonic/gin"
)

func (s *Server) DefineAuthRoutes(router *gin.RouterGroup) {

	repo := repositories.NewUserRepo(&s.DB)
	tokenRepo := repositories.NewTokenRepo(&s.DB)
	service := auth.NewAuthService(repo, tokenRepo)

	authRouter := router.Group("auth")

	authRouter.POST("login", service.Login)
	authRouter.POST("logout", service.Logout)
	authRouter.POST("register", service.Register)
	authRouter.POST("verify", service.Verify)
	authRouter.GET("me", service.GetMe)
}
