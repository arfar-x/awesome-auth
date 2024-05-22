package auth

import (
	"net/http"
	"strings"
	"time"

	"awesome-auth/configs"
	formrequest "awesome-auth/internal/core/http/request"
	"awesome-auth/internal/core/http/resources"
	"awesome-auth/internal/entities"
	"awesome-auth/internal/repositories"
	"awesome-auth/pkg/jwt"
	"awesome-auth/pkg/password"
	"awesome-auth/pkg/response"
	"github.com/gin-gonic/gin"
)

type Service struct {
	UserRepo  repositories.UserRepoInterface
	TokenRepo repositories.TokenRepoInterface
}

func NewAuthService(repo repositories.UserRepoInterface, tokenRepo repositories.TokenRepoInterface) *Service {
	return &Service{
		UserRepo:  repo,
		TokenRepo: tokenRepo,
	}
}

// Login user with given credentials.
func (srv *Service) Login(ctx *gin.Context) {
	defer recoverPanics(ctx, "Operation failed.")

	var request formrequest.LoginRequest
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, _ := srv.UserRepo.Get(ctx, entities.User{Username: request.Username})
	if user.ID == 0 {
		ctx.JSON(response.Unauthorized("Wrong credentials.", nil))
		return
	}

	tokenExpirationTime := time.Now().Add(time.Second * time.Duration(configs.Config.Jwt.ExpirationSeconds))
	tokenInstance := jwt.CreateToken(user.Username, tokenExpirationTime)
	token, _ := srv.TokenRepo.Create(ctx, entities.Token{
		UserID:    user.ID,
		Value:     tokenInstance.Value,
		ExpiresAt: tokenInstance.ExpiresAt,
	})

	if password.Check(user.Password, request.Password) {
		ctx.JSON(http.StatusOK, response.JsonResponse(
			"Logged in successfully.",
			http.StatusOK,
			map[string]string{
				"token":      token.Value,
				"expires_at": tokenInstance.ExpiresAt.Format("2006-01-02 15:04:05"),
			},
			nil,
		))
	} else {
		ctx.JSON(response.Unauthorized("Wrong credentials.", nil))
	}
}

// Logout user and expire the token.
func (srv *Service) Logout(ctx *gin.Context) {
	//defer recoverPanics(ctx, "")
	tokenString, _ := parseToken(ctx.GetHeader("Authorization"))

	tokenInstance := jwt.ParsePayload(tokenString)

	user, _ := srv.UserRepo.Get(ctx, entities.User{Username: tokenInstance.Payload})
	if user.ID == 0 {
		ctx.JSON(response.Unauthorized("Unauthorized.", nil))
		return
	}

	result, _ := srv.TokenRepo.Delete(ctx, entities.Token{
		Value:  tokenInstance.Value,
		UserID: user.ID,
	})
	if result {
		ctx.JSON(response.Success("Logged out.", result))
		return
	}

	ctx.JSON(response.NotFound("Could not log out.", result))
}

// Register user.
func (srv *Service) Register(ctx *gin.Context) {
	defer recoverPanics(ctx, "Could not create user.")

	var request formrequest.RegisterRequest
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, _ := srv.UserRepo.Create(ctx, entities.User{
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
	defer recoverPanics(ctx, "Unauthorized.")
	tokenString, _ := parseToken(ctx.GetHeader("Authorization"))

	tokenInstance := jwt.ParsePayload(tokenString)

	user, _ := srv.UserRepo.Get(ctx, entities.User{Username: tokenInstance.Payload})
	if user.ID == 0 {
		ctx.JSON(response.Unauthorized("Unauthorized.", nil))
		return
	}

	token, _ := srv.TokenRepo.Get(ctx, entities.Token{
		Value:  tokenInstance.Value,
		UserID: user.ID,
	})

	if token.ID != 0 && tokenInstance.Check() {
		ctx.JSON(response.Success("Valid.", true))
		return
	}

	ctx.JSON(response.Unauthorized("Unauthorized.", nil))
}

// GetMe gets the information about current user making the request.
func (srv *Service) GetMe(ctx *gin.Context) {
	defer recoverPanics(ctx, "")

	tokenString, _ := parseToken(ctx.GetHeader("Authorization"))

	tokenInstance := jwt.ParsePayload(tokenString)

	user, _ := srv.UserRepo.Get(ctx, entities.User{Username: tokenInstance.Payload})
	if user.ID == 0 {
		ctx.JSON(response.NotFound("User not found", nil))
		return
	}

	token, _ := srv.TokenRepo.Get(ctx, entities.Token{
		Value:  tokenInstance.Value,
		UserID: user.ID,
	})

	if token.ID == 0. || !tokenInstance.Check() {
		ctx.JSON(response.Unauthorized("Unauthorized.", nil))
		return
	}

	ctx.JSON(response.Success("User information.", resources.UserShowResource(user)))
}

func recoverPanics(ctx *gin.Context, message string) {
	if message == "" {
		message = "Operation failed."
	}
	if r := recover(); r != nil {
		ctx.AbortWithStatusJSON(response.InternalError(message, nil))
	}
}

func parseToken(token string) (value string, tokenType string) {
	splits := strings.Split(strings.TrimSpace(token), " ")
	value = splits[1]
	tokenType = strings.ToLower(splits[0])
	return
}
