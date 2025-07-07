package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType int

const (
	Income TransactionType = iota
	Expense
)

func (t TransactionType) String() string {
	return []string{"income", "expense"}[t]
}

type Transaction struct {
	ID          uuid.UUID       `json:"id"`
	UserID      uuid.UUID       `json:"user_id"`
	CategoryID  uuid.UUID       `json:"category_id"`
	Description string          `json:"description"`
	Amount      int64           `json:"amount"`
	Type        TransactionType `json:"type"`
	Date        time.Time       `json:"date"`
	CreatedAt   time.Time       `json:"created_at"`
}
