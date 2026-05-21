package util

import (
	"errors"
	"time"

	"github.com/MuhAndriJP/ayo-api/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	AdminID int64 `json:"admin_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(adminID int64) (string, error) {
	claims := Claims{
		AdminID: adminID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ayo-api",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.App.JWTSecret))
}

func ParseJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func keyFunc(t *jwt.Token) (any, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
	}
	return []byte(config.App.JWTSecret), nil
}
