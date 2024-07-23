package pgdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go-transaction-manager/internal/entity"
	"go-transaction-manager/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

var (
	ErrorNegativeBalance        error = errors.New("Balance can't be negative")
	ErrorUserNotFound           error = errors.New("User not found")
	ErrorNotEnoughMoney         error = errors.New("Not enough money")
	ErrorNotEnoughMoneyReserved error = errors.New("Not enough money reserved")
)

func NewUserRepository(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) AddBalance(ctx context.Context, id uuid.UUID, amount float64) error {
	const op = "repository.user.AddBalance"

	if amount < 0.0 {
		return fmt.Errorf("%s: %w", op, ErrorNegativeBalance)
	}

	err := r.DB.QueryRow(ctx, `INSERT INTO users_schema.user (id, balance, reserved) 
		VALUES($1, $2, 0) 
		ON CONFLICT (id) 
		DO UPDATE SET balance = users_schema.user.balance + $2`,
		id, amount).Scan()

	if err != nil {
		if err != pgx.ErrNoRows {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}

func (r *UserRepo) GetBalance(ctx context.Context, id uuid.UUID) (entity.Balance, error) {
	const op = "repository.user.GetBalance"

	var balance float64

	err := r.DB.QueryRow(ctx, `SELECT balance FROM users_schema.user  
		WHERE id = $1`,
		id).Scan(&balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Balance{}, fmt.Errorf("%s: %w", op, ErrorUserNotFound)
		}
		return entity.Balance{}, fmt.Errorf("%s: %w", op, err)
	}

	return entity.Balance{Balance: balance}, nil
}
