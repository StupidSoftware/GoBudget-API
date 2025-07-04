package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func initPostgres() (*pgxpool.Pool, error) {
	host := cfg.DBHost
	port := cfg.DBPort
	user := cfg.DBUser
	pass := cfg.DBPassword
	dbase := cfg.DBName

	psqlConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbase)
	poolConfig, err := pgxpool.ParseConfig(psqlConnStr)

	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}
