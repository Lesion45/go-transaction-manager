package service

import (
	"context"
	"go-transaction-manager/internal/entity"
	"go-transaction-manager/internal/repository"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) AddBalance(ctx context.Context, input UserAddBalanceInput) error {
	return s.userRepo.AddBalance(ctx, input.UserID, input.Amount)
}

func (s *UserService) GetBalance(ctx context.Context, input UserGetBalanceInput) (entity.Balance, error) {
	return s.userRepo.GetBalance(ctx, input.UserID)
}
