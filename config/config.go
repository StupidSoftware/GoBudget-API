package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db        *pgxpool.Pool
	logger    *Logger
	env       *conf
	v         *validator.Validate
	tokenAuth *jwt.Token
)

func Init() error {
	var err error

	env, err = InitEnv(".")
	if err != nil {
		return fmt.Errorf("error initializing config: %v", err)
	}

	db, err = initPostgres()
	if err != nil {
		return fmt.Errorf("error initializing postgres: %v", err)
	}

	v = NewValidator()
	tokenAuth = initJWT()

	return nil
}

func GetEnv() *conf {
	return env
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

func GetDB() *pgxpool.Pool {
	return db
}

func GetValidator() *validator.Validate {
	return v
}

func GetTokenAuth() *jwt.Token {
	return tokenAuth
}
