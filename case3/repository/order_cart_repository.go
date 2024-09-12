package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/elangreza14/qbit/case3/model"
)

type orderCartRepository struct {
	txRepo *PostgresTransactionRepo
}

func NewOrderCartRepository(
	tx PgxTXer,
) *orderCartRepository {
	return &orderCartRepository{
		txRepo: NewPostgresTransactionRepo(tx),
	}
}

func (ur *orderCartRepository) CreateOrderAndUpdateCart(ctx context.Context, order model.Order, cartIds []int) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		var err error
		orderRepo := NewOrderRepository(tx)
		err = orderRepo.Create(ctx, order)
		if err != nil {
			return err
		}

		cartRepo := NewCartRepository(tx)
		carts, err := cartRepo.GetChartByIDs(ctx, cartIds)
		if err != nil {
			return err
		}

		for _, cart := range carts {
			cart.UsedAt = sql.NullTime{Time: time.Now(), Valid: true}
			whereClause := make(map[string]any)
			whereClause["id"] = cart.ID
			err = cartRepo.Edit(ctx, cart, whereClause)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
