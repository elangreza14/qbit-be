package dto

import "github.com/google/uuid"

type (
	AddCartPayload struct {
		ProductID int `json:"product_id" binding:"required"`
		UserID    uuid.UUID
	}

	CartListResponseElement struct {
		ID           int    `json:"id"`
		Quantity     int    `json:"quantity"`
		Message      string `json:"message"`
		ProductID    int    `json:"product_id"`
		ProductName  string `json:"product_name"`
		ProductImage string `json:"product_image"`
		ProductPrice int    `json:"product_price"`
		ActualStock  int    `json:"actual_stock"`
	}

	CartListResponse []CartListResponseElement

	CheckoutCart struct {
		ChartIDs []int `json:"cart_ids" binding:"required,gt=0"`
		UserID   uuid.UUID
	}
)
