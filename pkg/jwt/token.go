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
	Claims    StandardClaimsWithUsername
	Instance  *jwt.Token
}

type StandardClaimsWithUsername struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
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
func CreateToken(claims StandardClaimsWithUsername, expiresAt time.Time) *Token {
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expiresAt.Unix(),
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
		Instance:  token,
	}
}

// ParsePayload parses the given token string and returns a Token instance initialized by parsed values.
func ParsePayload(tokenString string) *Token {
	token, err := jwt.ParseWithClaims(tokenString, &StandardClaimsWithUsername{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token algorithm is not valid")
		}
		return []byte(configs.Config.Jwt.SecretKey), nil
	})
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(*StandardClaimsWithUsername); ok && token.Valid {
		return &Token{
			Value:     tokenString,
			ExpiresAt: time.UnixMilli(claims.ExpiresAt),
			Claims:    *claims,
			Instance:  token,
		}
	}

	return nil
}
