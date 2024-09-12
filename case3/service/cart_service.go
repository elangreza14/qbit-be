package service

import (
	"context"
	"errors"

	"github.com/elangreza14/qbit/case3/dto"
	"github.com/elangreza14/qbit/case3/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type (
	cartRepo interface {
		GetAll(ctx context.Context) ([]model.Cart, error)
		Get(ctx context.Context, by string, val any, columns ...string) (*model.Cart, error)
		GetChartByUserIDAndProductID(ctx context.Context, userId uuid.UUID, productId int) (*model.Cart, error)
		Create(ctx context.Context, payloads ...model.Cart) error
		Edit(ctx context.Context, payload model.Cart, whereValues map[string]any) error
	}

	cartService struct {
		cartRepo    cartRepo
		productRepo productRepo
	}
)

func NewCartService(cartRepo cartRepo, productRepo productRepo) *cartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (cs *cartService) AddProductToCartList(ctx context.Context, req dto.AddCartPayload) error {

	// check product by id if more than one just add
	product, err := cs.productRepo.Get(ctx, "id", req.ProductID, "id", "stock")
	if err != nil {
		return err
	}

	if product.Stock < 1 {
		return errors.New("product is empty")
	}

	// check current cart exist or not

	cart, err := cs.cartRepo.GetChartByUserIDAndProductID(ctx, req.UserID, req.ProductID)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	// update or create
	if cart == nil {
		cart = model.NewCart(req.UserID, req.ProductID)
		err = cs.cartRepo.Create(ctx, *cart)
		if err != nil {
			return err
		}
	} else {

		cart.Quantity++

		where := make(map[string]any)
		where["product_id"] = cart.ProductID
		where["user_id"] = cart.UserID

		err = cs.cartRepo.Edit(ctx, *cart, where)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs *cartService) CartList(ctx context.Context) error {

	return nil
}
