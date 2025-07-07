package model

import "github.com/google/uuid"

type Category struct {
	ID     uuid.UUID  `json:"id"`
	Name   string     `json:"name" validate:"required,min=3,max=46" binding:"required"`
	UserID *uuid.UUID `json:"user_id"`
}
