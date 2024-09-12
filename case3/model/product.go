package model

import (
	"database/sql"
	"time"
)

type Product struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	Manufacturer string `db:"manufacturer"`
	Price        int    `db:"price"`
	Image        string `db:"image"`
	Stock        int    `db:"stock"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c Product) TableName() string {
	return "products"
}

// to create in DB
func (c Product) Data() map[string]any {
	return map[string]any{
		"id":           c.ID,
		"name":         c.Name,
		"manufacturer": c.Manufacturer,
		"price":        c.Price,
		"image":        c.Image,
		"stock":        c.Stock,
	}
}

func (c Product) Columns() []string {
	return []string{
		"id",
		"name",
		"manufacturer",
		"price",
		"image",
		"stock",
	}
}
