package service

import (
	"context"

	"github.com/elangreza14/qbit/case3/dto"
	"github.com/elangreza14/qbit/case3/model"
)

type (
	productsRepo interface {
		GetAll(ctx context.Context) ([]model.Product, error)
	}

	productsService struct {
		productsRepo productsRepo
	}
)

func NewProductService(productsRepo productsRepo) *productsService {
	return &productsService{
		productsRepo: productsRepo,
	}
}

func (cs *productsService) ProductsList(ctx context.Context) (dto.ProductListResponse, error) {
	products, err := cs.productsRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ProductListResponseElement, 0)
	for _, product := range products {
		res = append(res, dto.ProductListResponseElement{
			ID:           product.ID,
			DeviceName:   product.DeviceName,
			Manufacturer: product.Manufacturer,
			Price:        product.Price,
			Image:        product.Image,
			Stock:        product.Stock,
		})
	}

	return res, nil
}
