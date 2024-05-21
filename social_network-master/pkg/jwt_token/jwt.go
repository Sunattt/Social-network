package jwt_token

import (
	"Sunat/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func CreateToken(nik string, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nik_name": nik,
		"expire":   time.Now().Add(time.Minute * 10),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidToken(tokenStr string, secretKey string) (nikName string, valid bool, err error) {
	if tokenStr == "" {
		log.Println("Warning token emty!!")
		return "", false, err
	}
	claims := &models.TokenClaims{}
	// Парсим токен с использованием секретного ключа
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Println("ERROR validation token !!!")
		return "", false, err
	}

	if claims.Expire.Before(time.Now()) {
		log.Println("ERROR expire time")
		valid = false
		return "", valid, err
	} else {
		valid = token.Valid
		log.Println(valid)
		valid = true
	}

	log.Println(valid)
	nikName = claims.NikName
	return
}
