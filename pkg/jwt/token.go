package jwt

import (
	"fmt"
	"time"

	"awesome-auth/configs"
	"github.com/golang-jwt/jwt"
)

type Token struct {
	Value     string
	ExpiresAt time.Time
	Claims    jwt.MapClaims
	Payload   string
	Instance  *jwt.Token
}

// Check applies the conditions for a token to be valid. E.g. it makes sure token's expiration time does not
// exceed its valid time.
func (t *Token) Check() bool {
	if t.Instance.Valid {
		if !t.ExpiresAt.After(time.Now()) {
			return true
		}
	}

	return false
}

// CreateToken creates a Token instance by given payload and expiration time, also generates the token
// value and initializes Token by this value.
func CreateToken(payload any, expiresAt time.Time) *Token {
	claims := jwt.MapClaims{
		"payload":    payload,
		"expires_at": expiresAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configs.Config.Jwt.SecretKey))
	if err != nil {
		panic(err)
	}

	return &Token{
		Value:     tokenString,
		ExpiresAt: expiresAt,
		Claims:    claims,
		Payload:   payload.(string),
		Instance:  token,
	}
}

// ParsePayload parses the given token string and returns a Token instance initialized by parsed values.
func ParsePayload(tokenString string) *Token {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token algorithm is not valid")
		}
		return []byte(configs.Config.Jwt.SecretKey), nil
	})
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &Token{
			Value:     tokenString,
			ExpiresAt: time.UnixMilli(int64(claims["expires_at"].(float64))),
			Claims:    claims,
			Payload:   claims["payload"].(string),
			Instance:  token,
		}
	}

	return nil
}
