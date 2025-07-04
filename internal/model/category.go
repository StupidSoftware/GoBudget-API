package model

import "github.com/google/uuid"

type Category struct {
	ID     int
	Name   string
	UserID uuid.UUID
}
