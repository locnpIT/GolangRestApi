package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "asdkfasdjffhfaljkwehfagsdhgckawytue123y47yutqwgbefh!#!$#%#$%Y^rgkjhsddkf"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // token valid 2 hours
	})

	return token.SignedString([]byte(secretKey))
}
