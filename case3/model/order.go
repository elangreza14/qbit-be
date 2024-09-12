package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type (
	StatusEnum string

	Order struct {
		ID     uuid.UUID `db:"id"`
		UserID uuid.UUID `db:"user_id"`

		// Status can be
		// PROCESSING,
		// WAITING_PAYMENT,
		// SUCCESS,
		// FAILED
		Status StatusEnum `db:"status"`
		Total  int        `db:"total"`

		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
	}
)

const (
	ProcessingStatus     StatusEnum = "PROCESSING"
	WaitingPaymentStatus StatusEnum = "WAITING_PAYMENT"
	SuccessStatus        StatusEnum = "SUCCESS"
	FailedStatus         StatusEnum = "FAILED"
)

func NewOrder(ID uuid.UUID, userID uuid.UUID, total int) *Order {
	return &Order{
		ID:     ID,
		UserID: userID,
		Status: ProcessingStatus,
		Total:  total,
	}
}

func (u Order) TableName() string {
	return "orders"
}

// to create in DB
func (u Order) Data() map[string]any {
	return map[string]any{
		"id":      u.ID,
		"user_id": u.UserID,
		"status":  u.Status,
		"total":   u.Total,
	}
}

func (u Order) Columns() []string {
	return []string{
		"id",
		"user_id",
		"status",
		"total",
	}
}
