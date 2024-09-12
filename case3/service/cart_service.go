package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

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
		CheckAvailabilityCartList(ctx context.Context, userID uuid.UUID) ([]model.Cart, error)
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

	// check stock of cart
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

		if cart.Quantity > product.Stock {
			return errors.New("cannot add product, limited stock")
		}

		whereClause := make(map[string]any)
		whereClause["product_id"] = cart.ProductID
		whereClause["user_id"] = cart.UserID

		err = cs.cartRepo.Edit(ctx, *cart, whereClause)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs *cartService) CheckAvailabilityCartList(ctx context.Context, userID uuid.UUID) (dto.CartListResponse, error) {
	// check current carts exist or not
	carts, err := cs.cartRepo.CheckAvailabilityCartList(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := []dto.CartListResponseElement{}

	for _, cart := range carts {

		messageCode := cs.checkQuantityWithActualStock(cart.Quantity, cart.ActualStock)

		message := fmt.Sprintf("stock %s for this product %s", messageCode, cart.ProductName)

		res = append(res, dto.CartListResponseElement{
			ID:           cart.ID,
			Quantity:     cart.Quantity,
			Message:      message,
			ProductID:    cart.ProductID,
			ProductName:  cart.ProductName,
			ProductImage: cart.ProductImage,
			ProductPrice: cart.ProductPrice,
			ActualStock:  cart.ActualStock,
		})
	}

	return res, nil
}

func (cs *cartService) CheckoutSelectedProductsInCart(ctx context.Context, req dto.CheckoutCart) error {
	// check current carts exist or not
	carts, err := cs.cartRepo.CheckAvailabilityCartList(ctx, req.UserID)
	if err != nil {
		return err
	}

	unprocessedCartID := []string{}
	processedCartID := []int{}
	totalPrice := 0
	for _, cart := range carts {
		for _, chartID := range req.ChartIDs {
			if cart.ID == chartID {
				messageCode := cs.checkQuantityWithActualStock(cart.Quantity, cart.ActualStock)
				if messageCode != "AVAILABLE" {
					// return fmt.Errorf("stock %s for this product %s", messageCode, cart.ProductName)
					errMessage := fmt.Sprintf("stock %s for this product %s", messageCode, cart.ProductName)
					unprocessedCartID = append(unprocessedCartID, errMessage)
				}
			}

			processedCartID = append(processedCartID, chartID)
			totalPrice += cart.ProductPrice * cart.Quantity
		}
	}

	if len(unprocessedCartID) > 0 {
		return errors.New(strings.Join(unprocessedCartID, ", "))
	}

	if len(processedCartID) > 0 {
		orderID := uuid.New()
		fmt.Println(orderID, totalPrice)
	}

	return nil
}

func (cs *cartService) checkQuantityWithActualStock(quantity, actualStock int) string {
	if actualStock == 0 {
		return "NOT_AVAILABLE"
	}
	if quantity > actualStock {
		return "NOT_ENOUGH"
	}

	return "AVAILABLE"
}
