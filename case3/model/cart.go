package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        int       `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	ProductID int       `db:"product_id"`
	Quantity  int       `db:"quantity"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func NewCart(userID uuid.UUID, productID int) *Cart {
	return &Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  1,
	}
}

func (u Cart) TableName() string {
	return "carts"
}

// to create in DB
func (u Cart) Data() map[string]any {
	return map[string]any{
		"user_id":    u.UserID,
		"product_id": u.ProductID,
		"quantity":   u.Quantity,
	}
}
