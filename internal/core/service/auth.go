package service

import (
	"awesome-auth/internal/mysql"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	Repo mysql.RepoInterface
}

func NewAuthService(repo mysql.RepoInterface) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (srv *AuthService) Login(ctx *gin.Context) {

}

func (srv *AuthService) Logout(ctx *gin.Context) {

}

func (srv *AuthService) Register(ctx *gin.Context) {

}

func (srv *AuthService) Verify(ctx *gin.Context) {

}

func (srv *AuthService) GetMe(ctx *gin.Context) {

}
