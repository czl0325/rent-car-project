package token

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func GenerateToken(data interface{}) (string, error) {
	b, err := os.ReadFile("config/private.key")
	if err != nil {
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return "", err
	}
	now := time.Now().Unix() + 7200
	(data.(map[string]interface{}))["exp"] = now
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, data.(jwt.MapClaims))
	return token.SignedString(key)
}
