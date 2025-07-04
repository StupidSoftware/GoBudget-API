package server

import (
	"time"

	"github.com/breno5g/GoBudget/config"
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/google/uuid"
)

func Execute() {
	err := config.Init()
	logger := config.GetLogger("main")
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		panic(err)
	}

	var transaction model.Transaction

	transaction.ID = uuid.New()
	transaction.UserID = uuid.New()
	transaction.CategoryID = uuid.New()
	transaction.Description = "Test"
	transaction.Amount = 100
	transaction.Type = model.Income
	transaction.Date = time.Now()
	transaction.CreatedAt = time.Now()

	logger.Debug(transaction)

}
