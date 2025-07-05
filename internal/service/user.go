package service

import (
	"errors"
	"time"

	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/repository"
	"github.com/breno5g/GoBudget/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrClientNotFound = errors.New("cliente not found")
)

type UserService interface {
	Create(ctx *gin.Context, user *model.User) *utils.CustomError
	GetByUsername(ctx *gin.Context, username string) (*model.User, error)
	Delete(ctx *gin.Context, id string) error
}

type service struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx *gin.Context, user *model.User) *utils.CustomError {
	if err := user.Validate(); err != nil {
		return &utils.CustomError{
			Message: err[0].Message,
			Code:    400,
			Err:     errors.New(err[0].Message),
		}
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.HashPassword(user.Password)

	if err := s.repo.Create(ctx, user); err != nil {
		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) {
			if pgxErr.Code == "23505" {
				return &utils.CustomError{
					Message: "User already exists",
					Code:    400,
					Err:     errors.New("username already exists"),
				}
			}
		}

		return &utils.CustomError{
			Message: err.Error(),
			Code:    500,
			Err:     err,
		}
	}

	return nil
}

func (s *service) GetByUsername(ctx *gin.Context, username string) (*model.User, error) {
	return s.repo.GetByUsername(ctx, username)
}

func (s *service) Delete(ctx *gin.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
