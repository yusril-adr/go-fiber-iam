package utils

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	appErrors "iam-service/infrastructure/errors"
	"iam-service/modules/iam/auth/messages"
)

func CreateToken[T any](
	payload T,
	secret string,
	expiresIn int64,
) string {
	var payloadByte, err = json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": string(payloadByte),
		"exp":     expiresIn,
		"iat":     time.Now().Unix(),
		"nbf":     time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	secretInBytes := []byte(secret)
	tokenString, err := token.SignedString(secretInBytes)

	if err != nil {
		panic(err)
	}

	return tokenString
}

func ParseToken[T any](tokenString string, secret string) T {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			panic(appErrors.Unauthorized(messages.ErrTokenInvalidOrExpired))
		}

		if errors.Is(err, jwt.ErrTokenExpired) {
			panic(appErrors.Unauthorized(messages.ErrTokenInvalidOrExpired))
		}

		if errors.Is(err, jwt.ErrTokenMalformed) {
			panic(appErrors.Unauthorized(messages.ErrTokenInvalidOrExpired))
		}
		panic(err)
	}

	claims := token.Claims.(jwt.MapClaims)
	result := new(T)
	err = json.Unmarshal([]byte(claims["payload"].(string)), &result)
	if err != nil {
		panic(err)
	}
	return *result
}
