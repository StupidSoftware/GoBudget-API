package repository

import (
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository interface {
	Create(ctx *gin.Context, transaction *model.Transaction) error
	// GetByUserID(ctx *gin.Context, userID string) ([]*model.Transaction, error)
	// Update(ctx *gin.Context, transaction *model.Transaction) error
	// Delete(ctx *gin.Context, id string) error
}

type transactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) Create(ctx *gin.Context, transaction *model.Transaction) error {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	var dateValue interface{}
	if transaction.Date.Time.IsZero() {
		dateValue = nil
	} else {
		dateValue = transaction.Date.Time.Format("2006-01-02")
	}

	_, err = tx.Exec(ctx, "INSERT INTO transactions (id, user_id, category_id, description, amount, type, date) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		transaction.ID, transaction.UserID, transaction.CategoryID, transaction.Description, transaction.Amount, transaction.Type, dateValue)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
