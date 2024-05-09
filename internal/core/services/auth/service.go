package auth

import (
	"net/http"

	formrequest "awesome-auth/internal/core/http/request"
	"awesome-auth/internal/domain"
	"awesome-auth/internal/repositories"
	"awesome-auth/pkg/response"
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

// Register user.
func (srv *Service) Register(ctx *gin.Context) {
	var request formrequest.RegisterRequest
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := srv.Repo.Create(ctx, domain.UserDomain{
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password: func() string {
			hash, _ := password.Make(request.Password)
			return hash
		}(),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(response.InternalError("Could not create user.", nil))
		return
	}

	ctx.JSON(response.Created("User created successfully", resources.UserShowResource(result)))
}

func (srv *Service) Verify(ctx *gin.Context) {

}

func (srv *Service) GetMe(ctx *gin.Context) {

}
