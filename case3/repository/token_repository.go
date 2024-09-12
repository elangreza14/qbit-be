package repository

import (
	"github.com/elangreza14/qbit/case3/model"
)

type tokenRepository struct {
	db QueryPgx
	*PostgresRepo[model.Token]
}

func NewTokenRepository(
	dbPool QueryPgx,
) *tokenRepository {
	return &tokenRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Token](dbPool),
	}
}

func (*tokenRepository) NewTokenRepository(
	dbPool QueryPgx,
) *tokenRepository {
	return &tokenRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Token](dbPool),
	}
}
