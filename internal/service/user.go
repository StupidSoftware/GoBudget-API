package service

import (
	"context"
	"errors"
	"time"

	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/repository"
	"github.com/google/uuid"
)

var (
	ErrClientNotFound = errors.New("cliente not found")
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, user *model.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.HashPassword(user.Password)

	err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.repo.GetByUsername(ctx, username)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
