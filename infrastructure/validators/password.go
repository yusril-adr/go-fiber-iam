package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	PASSWORD_KEY     = "password"
	PASSWORD_MESSAGE = "must contain at least 1 each of Uppercase, Lowercase, Number and Special characters"
)

var (
	upperCaseRegex = regexp.MustCompile(`[A-Z]`)
	numberRegex    = regexp.MustCompile(`[0-9]`)
	specialRegex   = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	return upperCaseRegex.MatchString(password) &&
		numberRegex.MatchString(password) &&
		specialRegex.MatchString(password)
}
