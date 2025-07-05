package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username" validate:"required,min=3,max=20"`
	Password  string    `json:"password" validate:"required,min=12,max=144"`
	CreatedAt time.Time `json:"created_at"`
	Balance   int64     `json:"balance"`
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func formatValidationError(fe validator.FieldError) ValidationErrorResponse {
	var msg string

	switch fe.Tag() {
	case "required":
		msg = fe.Param() + " is required"
	case "min":
		msg = fe.Field() + " to be at least " + fe.Param() + " characters long"
	case "max":
		msg = fe.Field() + " to be at most " + fe.Param() + " characters long"
	default:
		msg = fe.Error()
	}

	return ValidationErrorResponse{
		Field:   fe.Field(),
		Message: msg,
	}
}

func (u *User) Validate() []ValidationErrorResponse {
	validate := validator.New()
	err := validate.Struct(u)

	if err == nil {
		return nil
	}

	var errs []ValidationErrorResponse
	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, formatValidationError(e))
	}

	return errs
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) HashPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(hashedPassword)
}
