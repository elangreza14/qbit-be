package controller

import (
	"context"
	"net/http"

	"github.com/elangreza14/qbit/case3/dto"
	"github.com/gin-gonic/gin"
)

type (
	productService interface {
		ProductList(ctx context.Context) (dto.ProductListResponse, error)
	}

	ProductController struct {
		productService
	}
)

func NewProductController(productService productService) *ProductController {
	return &ProductController{productService}
}

func (cc *ProductController) ProductList() gin.HandlerFunc {
	return func(c *gin.Context) {
		currencies, err := cc.productService.ProductList(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(currencies, nil))
	}
}
