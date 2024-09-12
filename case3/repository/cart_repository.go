package repository

import (
	"context"
	"fmt"

	"github.com/elangreza14/qbit/case3/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

type cartRepository struct {
	db QueryPgx
	*PostgresRepo[model.Cart]
}

func NewCartRepository(
	dbPool QueryPgx,
) *cartRepository {
	return &cartRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Cart](dbPool),
	}
}

func (cr *cartRepository) GetChartByUserIDAndProductID(ctx context.Context, userId uuid.UUID, productId int) (*model.Cart, error) {
	q := fmt.Sprintf(cr.QueryBasicSelect + ` WHERE user_id = $1 AND product_id = $2 LIMIT 1`)

	v, err := pgxutil.SelectRow(ctx, cr.db, q, []any{userId, productId}, pgx.RowToAddrOfStructByNameLax[model.Cart])
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (cr *cartRepository) CheckAvailabilityCartList(ctx context.Context, userID uuid.UUID) ([]model.Cart, error) {
	q := `SELECT 
			c.id,
			c.user_id,
			c.product_id,
			c.quantity,
			c.created_at,
			c.updated_at,
			p.stock AS actual_stock,
			p.image AS product_image,
			p.name AS product_name,
			p.price AS product_price
		FROM carts c LEFT JOIN products p ON c.product_id = p.id WHERE user_id = $1`

	v, err := pgxutil.Select(ctx, cr.db, q, []any{userID}, pgx.RowToStructByNameLax[model.Cart])
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (cr *cartRepository) GetCartByIDs(ctx context.Context, charIds []int) ([]model.Cart, error) {
	q := `SELECT 
			c.id,
			c.user_id,
			c.product_id,
			c.quantity,
			c.created_at,
			c.updated_at,
			p.stock AS actual_stock,
			p.image AS product_image,
			p.name AS product_name,
			p.price AS product_price
		FROM carts c LEFT JOIN products p ON c.product_id = p.id WHERE c.id = ANY($1)`

	v, err := pgxutil.Select(ctx, cr.db, q, []any{charIds}, pgx.RowToStructByNameLax[model.Cart])
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (cr *cartRepository) DeleteCartByIDs(ctx context.Context, id int) error {
	q := `DELETE FROM carts WHERE id=$1;`

	_, err := cr.db.Exec(ctx, q, id)
	return err
}
