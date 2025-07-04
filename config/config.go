package config

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db     *pgxpool.Pool
	logger *Logger
	cfg    *conf
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
