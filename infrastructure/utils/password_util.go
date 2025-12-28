package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	if err != nil {
		panic(err.Error())
	}

	return string(b)
}

func CheckPassword(pw, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}
