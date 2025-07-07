package service

import (
	"errors"
	"time"

	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/repository"
	"github.com/breno5g/GoBudget/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		return utils.NewCustomPGError("user already exists", 409, err)
	}

	return nil
}

func (s *userService) Login(ctx *gin.Context, user *model.User) (string, *utils.CustomError) {
	dbUser, err := s.repo.GetByUsername(ctx, user.Username)

	userNotFound := utils.NewCustomPGError("User not found", 404, errors.New("user not found"))

	if err != nil {
		userNotFound.Err = err
		return "", userNotFound
	}

	if !user.ComparePassword(dbUser.Password) {
		userNotFound.Err = err
		return "", userNotFound
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
