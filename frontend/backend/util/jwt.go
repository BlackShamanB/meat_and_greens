package util

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

const SecretKey = "secret"

func GenerateJwt(issure string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issure,
		ExpiresAt: jwt.NewTime(float64(time.Now().Add(time.Hour * 24).Unix())),
	})

	return claims.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}
