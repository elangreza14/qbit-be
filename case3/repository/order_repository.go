package repository

import (
	"github.com/elangreza14/qbit/case3/model"
)

type orderRepository struct {
	db QueryPgx
	*PostgresRepo[model.Order]
}

func NewOrderRepository(
	dbPool QueryPgx,
) *orderRepository {
	return &orderRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Order](dbPool),
	}
}
