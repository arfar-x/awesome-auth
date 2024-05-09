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

func CreateToken(payload string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":   payload,
			"expires_at": time.Now().Add(time.Second * time.Duration(configs.Config.Jwt.ExpirationSeconds)).Unix(),
		})

	tokenString, err := token.SignedString([]byte(configs.Config.Jwt.SecretKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func Validate(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Config.Jwt.SecretKey), nil
	})

	if err != nil {
		panic(err)
	}

	return token.Valid
}
