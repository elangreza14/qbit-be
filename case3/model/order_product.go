package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type (
	OrderProduct struct {
		ID          int       `db:"id"`
		OrderID     uuid.UUID `db:"order_id"`
		ProductID   int       `db:"product_id"`
		Quantity    int       `db:"quantity"`
		Price       int       `db:"price"`
		ActualStock int       `db:"actual_stock"`

		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
	}
)

func NewOrderProduct(orderID uuid.UUID, productID, quantity, price int) *OrderProduct {
	return &OrderProduct{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
}

func (u OrderProduct) TableName() string {
	return "order_products"
}

func (u OrderProduct) Data() map[string]any {
	return map[string]any{
		"order_id":   u.OrderID,
		"product_id": u.ProductID,
		"quantity":   u.Quantity,
		"price":      u.Price,
	}
}

func (u OrderProduct) Columns() []string {
	return []string{
		"id",
		"order_id",
		"product_id",
		"quantity",
		"price",
	}
}
