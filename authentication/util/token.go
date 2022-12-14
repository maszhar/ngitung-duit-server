package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyTokenClaims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func CreateAccessToken(key string, userId string) (string, error) {
	expires := time.Now().Add(time.Hour)

	claims := MyTokenClaims{
		"ACCESS_TOKEN",
		jwt.RegisteredClaims{
			Subject:   userId,
			Issuer:    "authorizer",
			ExpiresAt: jwt.NewNumericDate(expires),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	payload, err := token.SignedString([]byte(key))
	return payload, err
}

func ParseAccessToken(key string, payload string) (string, error) {
	var invalidToken = errors.New("invalid token method")

	token, err := jwt.ParseWithClaims(payload, &MyTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidToken
		}
		return []byte(key), nil
	})
	if err != nil {
		return "", invalidToken
	}

	if claims, ok := token.Claims.(*MyTokenClaims); ok && token.Valid {
		if claims.Type != "ACCESS_TOKEN" {
			return "", invalidToken
		}
		return claims.Subject, nil
	} else {
		return "", invalidToken
	}
}
