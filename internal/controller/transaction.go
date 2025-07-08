package controller

import (
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/service"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	Create(ctx *gin.Context)
	GetByUserID(ctx *gin.Context)
}

type transactionController struct {
	svc service.TransactionService
}

func NewTransactionController(svc service.TransactionService) *transactionController {
	return &transactionController{
		svc: svc,
	}
}

func (t *transactionController) Create(ctx *gin.Context) {
	var transaction model.Transaction

	if err := ctx.BindJSON(&transaction); err != nil {
		logger.Errorf("Error binding JSON: %v", err)
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	logger.Debugf("Transaction: %+v", transaction)

	if err := t.svc.Create(ctx, &transaction); err != nil {
		logger.Errorf("Error creating transaction: %v", err)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Transaction created",
	})
}

func (t *transactionController) GetByUserID(ctx *gin.Context) {
	transactions, err := t.svc.GetByUserID(ctx)
	if err != nil {
		logger.Errorf("Error getting transactions: %v", err)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(200, transactions)
}
