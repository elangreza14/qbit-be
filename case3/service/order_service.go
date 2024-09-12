package service

import (
	"context"

	"github.com/elangreza14/qbit/case3/dto"
	"github.com/elangreza14/qbit/case3/model"
	"github.com/google/uuid"
)

type (
	orderRepo interface {
		GetAll(ctx context.Context) ([]model.Order, error)
		Create(ctx context.Context, payloads ...model.Order) error
		Edit(ctx context.Context, payload model.Order, whereValues map[string]any) error
	}

	orderCartProductRepository interface {
		UpdateOrderAndProduct(ctx context.Context, orderID uuid.UUID, cartIds []int) error
	}

	orderService struct {
		orderCartProductRepository orderCartProductRepository
	}
)

func NewOrderService(orderCartProductRepository orderCartProductRepository) *orderService {
	return &orderService{
		orderCartProductRepository: orderCartProductRepository,
	}
}

func (cs *orderService) UpdateOrder(ctx context.Context, req dto.UpdateOrder) error {
	return cs.orderCartProductRepository.UpdateOrderAndProduct(ctx, req.OrderID, req.CartIDs)
}
