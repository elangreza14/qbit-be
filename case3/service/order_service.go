package service

import (
	"context"
	"fmt"

	"github.com/elangreza14/qbit/case3/dto"
	"github.com/elangreza14/qbit/case3/model"
)

type (
	orderRepo interface {
		GetAll(ctx context.Context) ([]model.Order, error)
		Create(ctx context.Context, payloads ...model.Order) error
		Edit(ctx context.Context, payload model.Order, whereValues map[string]any) error
	}

	orderService struct {
		cartRepo  cartRepo
		orderRepo orderRepo
	}
)

func NewOrderService(cartRepo cartRepo, orderRepo orderRepo) *orderService {
	return &orderService{
		cartRepo:  cartRepo,
		orderRepo: orderRepo,
	}
}

func (cs *orderService) UpdateOrder(ctx context.Context, req dto.UpdateOrder) error {
	fmt.Println("bisa update sekarang")
	return nil
}
