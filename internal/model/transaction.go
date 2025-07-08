package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/breno5g/GoBudget/internal/utils"
	"github.com/google/uuid"
)

type TransactionType int

const (
	Income TransactionType = iota
	Expense
)

func (t TransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TransactionType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch strings.ToLower(str) {
	case "income":
		*t = Income
	case "expense":
		*t = Expense
	default:
		return fmt.Errorf("invalid transaction type: %s", str)
	}

	return nil
}

func (t TransactionType) String() string {
	return []string{"income", "expense"}[t]
}

type Transaction struct {
	ID          uuid.UUID        `json:"id"`
	UserID      uuid.UUID        `json:"user_id"`
	CategoryID  uuid.UUID        `json:"category_id"`
	Description string           `json:"description" validate:"required,min=3,max=255" binding:"required"`
	Amount      int64            `json:"amount" validate:"required,gt=0" binding:"required"`
	Type        *TransactionType `json:"type" validate:"required" binding:"required"`
	Date        utils.Date       `json:"date" validate:"required" binding:"required"`
	CreatedAt   time.Time        `json:"created_at"`
}
