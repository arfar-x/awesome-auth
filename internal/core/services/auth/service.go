package auth

import (
	"net/http"

	formrequest "awesome-auth/internal/core/http/request"
	"awesome-auth/internal/repositories"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Repo repositories.RepoInterface
}

func NewAuthService(repo repositories.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (srv *Service) Login(ctx *gin.Context) {

}

func (srv *Service) Logout(ctx *gin.Context) {

}

func (srv *Service) Register(ctx *gin.Context) {
	var request formrequest.RegisterRequest
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := srv.Repo.Create( /* WHAT TO SEND HERE ? */ )

	ctx.JSON(http.StatusOK, map[string]any{
		"data": result,
	})
}

func (srv *Service) Verify(ctx *gin.Context) {

}

func (srv *Service) GetMe(ctx *gin.Context) {

}
