package auth

import (
	"net/http"

	formrequest "awesome-auth/internal/core/http/request"
	"awesome-auth/internal/core/http/resources"
	"awesome-auth/internal/domain"
	"awesome-auth/internal/repositories"
	"awesome-auth/pkg/password"
	"awesome-auth/pkg/response"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Repo      repositories.RepoInterface
	TokenRepo repositories.RepoInterface
}

func NewAuthService(repo repositories.RepoInterface, tokenRepo repositories.RepoInterface) *Service {
	return &Service{
		Repo:      repo,
		TokenRepo: tokenRepo,
	}
}

// Login user with given credentials.
func (srv *Service) Login(ctx *gin.Context) {
	//defer recoverPanics(ctx, "Operation failed.")

	var request formrequest.LoginRequest
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, _ := srv.Repo.Get(ctx, domain.UserDomain{Username: request.Username})

	token := srv.Repo..CreateToken(ctx)

	if password.Check(user.Password, request.Password) {
		//token := jwt.CreateToken(user.Username)
		ctx.JSON(http.StatusOK, response.JsonResponse(
			"Logged in successfully.",
			http.StatusOK,
			map[string]string{"token": token},
			nil,
		))
	} else {
		ctx.JSON(response.Unauthorized("Wrong credentials.", nil))
	}
}

// Logout user and expire the token.
func (srv *Service) Logout(ctx *gin.Context) {

}

// Register user.
func (srv *Service) Register(ctx *gin.Context) {
	defer recoverPanics(ctx, "Could not create user.")

	var request formrequest.RegisterRequest
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, _ := srv.Repo.Create(ctx, domain.UserDomain{
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password: func() string {
			hash, _ := password.Make(request.Password)
			return hash
		}(),
	})

	ctx.JSON(response.Created("User created successfully", resources.UserShowResource(result)))
}

// Verify whether the user is authorized or not.
func (srv *Service) Verify(ctx *gin.Context) {

}

func (srv *Service) GetMe(ctx *gin.Context) {

}

func recoverPanics(ctx *gin.Context, message string) {
	if message == "" {
		message = "Operation failed."
	}
	if r := recover(); r != nil {
		ctx.AbortWithStatusJSON(response.InternalError(message, nil))
	}
}
