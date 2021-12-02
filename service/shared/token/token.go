package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type userTokenClaim struct {
	Phone string
	jwt.StandardClaims
}

func GenerateToken(phone string) (string, error) {
	b, err := os.ReadFile("config/private.key")
	if err != nil {
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return "", err
	}
	now := time.Now().Unix() + 7200
	claim := userTokenClaim{
		phone,
		jwt.StandardClaims{
			ExpiresAt: now,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claim)
	return token.SignedString(key)
}

func VerifyToken(token string) (interface{}, error) {
	b, err := os.ReadFile("config/public.key")
	if err != nil {
		return nil, err
	}
	result, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
		if err != nil {
			return nil, err
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token解析失败，错误=%q\n", err)
	}
	return result.Claims, nil
}