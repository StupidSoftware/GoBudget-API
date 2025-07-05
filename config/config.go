package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db     *pgxpool.Pool
	logger *Logger
	cfg    *conf
	v      *validator.Validate
)

func Init() error {
	var err error

	cfg, err = InitEnv(".")
	if err != nil {
		return fmt.Errorf("error initializing config: %v", err)
	}

	db, err = initPostgres()
	if err != nil {
		return fmt.Errorf("error initializing postgres: %v", err)
	}

	v = NewValidator()

	return nil
}

func GetEnv() *conf {
	return cfg
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
