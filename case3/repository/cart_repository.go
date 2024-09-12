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
