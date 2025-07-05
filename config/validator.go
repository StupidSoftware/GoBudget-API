package config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func validateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 12 || len(password) > 144 {
		return false
	}

	upperCasePattern := `[A-Z]`
	if matched, _ := regexp.MatchString(upperCasePattern, password); !matched {
		return false
	}

	specialCharPattern := `[!@#~$%^&*()_+|<>?:{}]`
	if matched, _ := regexp.MatchString(specialCharPattern, password); !matched {
		return false
	}

	return true
}

func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("str_pswd", validateStrongPassword)
	return v
}
