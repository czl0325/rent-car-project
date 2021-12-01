package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func main() {
	b, err := os.ReadFile("config/private.key")
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		panic(err)
	}
	now := time.Now().Unix() + 7200
	fmt.Printf("时间=%d\n", now)
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"exp":   now,
		"phone": "123",
	})
	s, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后的jwt字符串=%q\n", s)

	b, err = os.ReadFile("config/public.key")
	if err != nil {
		panic(err)
	}
	token2, err := jwt.ParseWithClaims(s, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
		if err != nil {
			return nil, err
		}
		return publicKey, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后的结果=%+v\n", token2.Claims)
}
