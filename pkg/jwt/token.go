package jwt

import (
	"time"

	"awesome-auth/configs"
	"github.com/golang-jwt/jwt"
)

type Token struct {
	value     string
	expiresAt time.Time
}

func (t *Token) NewToken(payload string) Token {
	return Token{value: payload}
}

func CreateToken(payload string, expiresAt time.Time) (string, time.Time) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"payload":    payload,
			"expires_at": expiresAt.Unix(),
		})

	tokenString, err := token.SignedString([]byte(configs.Config.Jwt.SecretKey))
	if err != nil {
		panic(err)
	}
	return tokenString, expiresAt
}

func Validate(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Config.Jwt.SecretKey), nil
	})

	if err != nil {
		panic(err)
	}

	// TODO: Check token expiration time

	return token.Valid
}
