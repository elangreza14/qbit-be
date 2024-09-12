package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/elangreza14/qbit/case3/dto"
	"github.com/elangreza14/qbit/case3/middleware"
	"github.com/elangreza14/qbit/case3/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	cartService interface {
		AddProductToCartList(ctx context.Context, req dto.AddCartPayload) error
		CheckAvailabilityCartList(ctx context.Context, userID uuid.UUID) (dto.CartListResponse, error)
		CheckoutSelectedProductsInCart(ctx context.Context, req dto.CheckoutCart) error
	}

	CartController struct {
		cartService
	}
)

func NewCartController(cartService cartService) *CartController {
	return &CartController{cartService}
}

func (cc *CartController) AddProductToCartList() gin.HandlerFunc {
	return func(c *gin.Context) {

		v, ok := c.Get(middleware.UserMiddlewareKey)
		if !ok {
			err := errors.New("error reading middleware")
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		user, ok := v.(*model.User)
		if !ok {
			err := errors.New("not valid middleware")
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		req := dto.AddCartPayload{}
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}
		req.UserID = user.ID

		err = cc.cartService.AddProductToCartList(c, req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(nil, nil))
	}
}

func (cc *CartController) CheckAvailabilityCartList() gin.HandlerFunc {
	return func(c *gin.Context) {

		v, ok := c.Get(middleware.UserMiddlewareKey)
		if !ok {
			err := errors.New("error reading middleware")
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		user, ok := v.(*model.User)
		if !ok {
			err := errors.New("not valid middleware")
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		carts, err := cc.cartService.CheckAvailabilityCartList(c, user.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(carts, nil))
	}
}

func (cc *CartController) CheckoutSelectedProductsInCart() gin.HandlerFunc {
	return func(c *gin.Context) {

		v, ok := c.Get(middleware.UserMiddlewareKey)
		if !ok {
			err := errors.New("error reading middleware")
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		user, ok := v.(*model.User)
		if !ok {
			err := errors.New("not valid middleware")
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		req := dto.CheckoutCart{}
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}
		req.UserID = user.ID

		err = cc.cartService.CheckoutSelectedProductsInCart(c, req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(nil, nil))
	}
}
