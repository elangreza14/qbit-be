package dto

import "github.com/google/uuid"

type (
	AddCartPayload struct {
		ProductID int `json:"product_id" binding:"required"`
		UserID    uuid.UUID
	}
)
