package repository

import (
	"context"

	"github.com/elangreza14/qbit/case3/model"
)

type productRepository struct {
	db QueryPgx
	*PostgresRepo[model.Product]
}

func NewProductRepository(
	dbPool QueryPgx,
) *productRepository {
	return &productRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Product](dbPool),
	}
}

func (cr *productRepository) UpdateStockByID(ctx context.Context, stock int, id int) error {
	_, err := cr.db.Exec(ctx, `UPDATE products SET stock=$1 WHERE id=$2;`, stock, id)
	return err
}
