package service

import (
	"errors"

	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/repository"
	"github.com/breno5g/GoBudget/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionService interface {
	Create(ctx *gin.Context, transaction *model.Transaction) *utils.CustomError
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *transactionService {
	return &transactionService{
		repo: repo,
	}
}

func (t *transactionService) Create(ctx *gin.Context, transaction *model.Transaction) *utils.CustomError {
	transaction.ID = uuid.New()
	userID, ok := ctx.Get("user_id")

	if !ok {
		return &utils.CustomError{
			Message: "user id is required",
			Code:    400,
			Err:     errors.New("user id is required"),
		}
	}

	transaction.UserID = uuid.MustParse(userID.(string))

	if err := t.repo.Create(ctx, transaction); err != nil {
		return utils.NewCustomPGError("error creating transaction", 500, err)
	}

	return nil
}
