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
	Login(ctx *gin.Context, user *model.User) (string, *utils.CustomError)
	Delete(ctx *gin.Context, id string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx *gin.Context, user *model.User) *utils.CustomError {
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

func (s *userService) Login(ctx *gin.Context, user *model.User) (string, *utils.CustomError) {
	dbUser, err := s.repo.GetByUsername(ctx, user.Username)

	if err != nil {
		return "", &utils.CustomError{
			Message: "User not found",
			Code:    404,
			Err:     errors.New("user not found"),
		}
	}

	if !user.ComparePassword(dbUser.Password) {
		return "", &utils.CustomError{
			Message: "User not found",
			Code:    404,
			Err:     errors.New("user not found"),
		}
	}

	token, err := utils.GenerateToken(dbUser.ID.String())

	if err != nil {
		return "", &utils.CustomError{
			Message: "Error generating token",
			Code:    500,
			Err:     err,
		}
	}

	return token, nil

}

func (s *userService) Delete(ctx *gin.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
