package repository

import (
	"context"

	"github.com/elangreza14/qbit/case3/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

type orderProductRepository struct {
	db QueryPgx
	*PostgresRepo[model.OrderProduct]
}

func NewOrderProductRepository(
	dbPool QueryPgx,
) *orderProductRepository {
	return &orderProductRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.OrderProduct](dbPool),
	}
}

func (cr *orderProductRepository) GetOrderProductByOrderID(ctx context.Context, orderID uuid.UUID) ([]model.OrderProduct, error) {
	q := `SELECT op.id, op.order_id, op.product_id, op.quantity, op.price, op.created_at, op.updated_at, p.stock as actual_stock 
			FROM order_products op LEFT JOIN products p ON p.id = op.product_id
		WHERE op.order_id = $1`

	v, err := pgxutil.Select(ctx, cr.db, q, []any{orderID}, pgx.RowToStructByNameLax[model.OrderProduct])
	if err != nil {
		return nil, err
	}
	return v, nil
}
