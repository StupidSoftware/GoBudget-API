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
	ID          uuid.UUID
	UserID      uuid.UUID
	CategoryID  uuid.UUID
	Description string
	Amount      int64
	Type        TransactionType
	Date        time.Time
	CreatedAt   time.Time
}
