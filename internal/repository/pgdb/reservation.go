package pgdb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"go-transaction-manager/internal/entity"
	"go-transaction-manager/pkg/postgres"
	"time"
)

type ReservationRepo struct {
	*postgres.Postgres
}

func NewReservationRepository(pg *postgres.Postgres) *ReservationRepo {
	return &ReservationRepo{pg}
}

func (r *ReservationRepo) ReserveBalance(ctx context.Context, reservation entity.Reservation) error {
	const op = "repository.reservation.ReserveBalance"

	var balance, reservedBalance float64

	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, `UPDATE users_schema.user 
		SET balance = balance - $2, reserved = reserved + $2
		WHERE id = $1
		RETURNING balance, reserved`,
		reservation.UserID, reservation.Amount).Scan(&balance, &reservedBalance)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if balance < 0 {
		return fmt.Errorf("%s: %s", op, ErrorNotEnoughMoney)
	}

	err = tx.QueryRow(ctx, `INSERT INTO orders_schema.order(id, user_id, service_id, amount, info) VALUES($1, $2, #3, $4, $5)`,
		reservation.OrderID, reservation.UserID, reservation.ServiceID, reservation.Amount, reservation.Info).Scan()
	if err != nil {
		if err != pgx.ErrNoRows {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *ReservationRepo) CommitReservedBalance(ctx context.Context, reservation entity.Reservation) error {
	const op = "repository.user.CommitReservedBalance"

	var reservedBalance float64

	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, `UPDATE users_schema.user 
		SET reserved = reserved - $2
		WHERE id = $1
		RETURNING reserved`,
		reservation.UserID, reservation.Amount).Scan(&reservedBalance)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if reservedBalance < 0 {
		return fmt.Errorf("%s: %s", op, ErrorNotEnoughMoneyReserved)
	}

	err = tx.QueryRow(context.Background(), `DELETE FROM orders_schema.order WHERE id = $1`, reservation.OrderID).Scan()
	if err != nil {
		if err != pgx.ErrNoRows {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	err = tx.QueryRow(ctx, `INSERT INTO reports_schema.report(date, service_id, user_id, amount, info) VALUES($1, $2, #3, $4, $5)`,
		time.DateOnly, reservation.UserID, reservation.ServiceID, reservation.Amount, reservation.Info).Scan()
	if err != nil {
		if err != pgx.ErrNoRows {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
