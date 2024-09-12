package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID           int       `db:"id"`
	UserID       uuid.UUID `db:"user_id"`
	ProductID    int       `db:"product_id"`
	Name         string    `db:"name"`
	CurrencyCode string    `db:"currency_code"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func NewAccount(userID uuid.UUID, name, currencyCode string, productID int) (*Account, error) {
	return &Account{
		UserID:       userID,
		ProductID:    productID,
		Name:         name,
		CurrencyCode: currencyCode,
	}, nil
}

func (a Account) TableName() string {
	return "accounts"
}

func (a Account) Columns() []string {
	return []string{
		"id",
		"user_id",
		"name",
		"currency_code",
		"product_id",
	}
}

// to create in DB
func (a Account) Data() map[string]any {
	return map[string]any{
		"id":            a.ID,
		"user_id":       a.UserID,
		"name":          a.Name,
		"currency_code": a.CurrencyCode,
		"product_id":    a.ProductID,
	}
}
