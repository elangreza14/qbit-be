package model

import (
	"database/sql"
	"time"
)

type Product struct {
	ID           int    `db:"id"`
	DeviceName   string `db:"device_name"`
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

func (c Product) Columns() []string {
	return []string{
		"id",
		"device_name",
		"manufacturer",
		"price",
		"image",
		"stock",
		"created_at",
		"updated_at",
	}
}

// to create in DB
func (c Product) Data() map[string]any {
	return map[string]any{
		"id":           c.ID,
		"device_name":  c.DeviceName,
		"manufacturer": c.Manufacturer,
		"price":        c.Price,
		"image":        c.Image,
		"stock":        c.Stock,
		"created_at":   c.CreatedAt,
		"updated_at":   c.UpdatedAt,
	}
}
