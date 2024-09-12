package service

import (
	"context"

	"github.com/elangreza14/qbit/case3/dto"
	"github.com/elangreza14/qbit/case3/model"
)

type (
	productRepo interface {
		GetAll(ctx context.Context) ([]model.Product, error)
		Get(ctx context.Context, by string, val any, columns ...string) (*model.Product, error)
	}

	productService struct {
		productRepo productRepo
	}
)

func NewProductService(productRepo productRepo) *productService {
	return &productService{
		productRepo: productRepo,
	}
}

func (cs *productService) ProductList(ctx context.Context) (dto.ProductListResponse, error) {
	product, err := cs.productRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ProductListResponseElement, 0)
	for _, product := range product {
		res = append(res, dto.ProductListResponseElement{
			ID:           product.ID,
			DeviceName:   product.Name,
			Manufacturer: product.Manufacturer,
			Price:        product.Price,
			Image:        product.Image,
			Stock:        product.Stock,
		})
	}

	return res, nil
}
