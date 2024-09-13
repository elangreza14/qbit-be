package repository

import (
	"context"
	"errors"

	"github.com/elangreza14/qbit/case3/model"
	"github.com/google/uuid"
)

type orderCartProductRepository struct {
	txRepo *PostgresTransactionRepo
}

func NewOrderCartProductRepository(
	tx PgxTXer,
) *orderCartProductRepository {
	return &orderCartProductRepository{
		txRepo: NewPostgresTransactionRepo(tx),
	}
}

func (ur *orderCartProductRepository) CreateOrderAndUpdateCart(ctx context.Context, userID uuid.UUID, cartIds []int) (uuid.UUID, error) {
	orderID := uuid.New()

	err := ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		var err error

		cartRepo := NewCartRepository(tx)
		carts, err := cartRepo.GetCartByIDs(ctx, cartIds)
		if err != nil {
			return err
		}

		orderProductRepo := NewOrderProductRepository(tx)

		totalPrice := 0
		orderProducts := make([]model.OrderProduct, 0)
		for _, cart := range carts {
			orderProducts = append(orderProducts, model.OrderProduct{
				OrderID:   orderID,
				ProductID: cart.ProductID,
				Quantity:  cart.Quantity,
				Price:     cart.ProductPrice,
			})
			totalPrice += cart.Quantity * cart.ProductPrice
			err = cartRepo.DeleteCartByIDs(ctx, cart.ID)
			if err != nil {
				return err
			}
		}

		order := model.NewOrder(orderID, userID, totalPrice)
		orderRepo := NewOrderRepository(tx)
		err = orderRepo.Create(ctx, *order)
		if err != nil {
			return err
		}

		err = orderProductRepo.Create(ctx, orderProducts...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	return orderID, nil
}

func (ur *orderCartProductRepository) UpdateOrderAndProduct(ctx context.Context, orderID uuid.UUID, cartIds []int) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		var err error
		orderProductRepo := NewOrderProductRepository(tx)
		orderProducts, err := orderProductRepo.GetOrderProductByOrderID(ctx, orderID)
		if err != nil {
			return err
		}

		orderRepo := NewOrderRepository(tx)
		order, err := orderRepo.Get(ctx, "id", orderID)
		if err != nil {
			return err
		}

		productRepo := NewProductRepository(tx)
		for _, orderProduct := range orderProducts {
			if orderProduct.ActualStock >= orderProduct.Quantity {
				stock := orderProduct.ActualStock - orderProduct.Quantity
				if stock < 0 {
					return errors.New("error when calculating stock")
				}

				err = productRepo.UpdateStockByID(ctx, stock, orderProduct.ProductID)
				if err != nil {
					return err
				}
			} else {
				order.Status = "CANCELED_BECAUSE_LIMITED_STOCK"
				whereClause := make(map[string]any)
				whereClause["id"] = order.ID
				err = orderRepo.Edit(ctx, *order, whereClause)
				if err != nil {
					return err
				}

				return nil
			}
		}

		order.Status = "WAITING_PAYMENT"
		whereClause := make(map[string]any)
		whereClause["id"] = order.ID
		err = orderRepo.Edit(ctx, *order, whereClause)
		if err != nil {
			return err
		}

		return nil
	})
}
