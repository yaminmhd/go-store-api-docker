package service

import (
	"context"
	"github.com/yaminmhd/go-hardware-store/contract"
	"github.com/yaminmhd/go-hardware-store/log"
)

type ProductService struct {
	productRepository ProductRepository
}

func (service ProductService) GetProducts(ctx context.Context) (contract.GetProducts, error) {
	products, err := service.productRepository.GetAllProducts(ctx)

	if err != nil {
		log.Log.Error("[ProductService] error getting products", err)
		return contract.GetProducts{}, err
	}

	var productResult []contract.ProductSummary

	for _, product := range products {
		productResult = append(productResult, contract.ProductSummary{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  product.Quantity,
			State:     product.State,
		})
	}

	productsList := contract.GetProducts{
		Products:   productResult,
		TotalCount: len(products),
	}
	return productsList, nil
}

func NewProductService(productRepository ProductRepository) ProductService {
	return ProductService{productRepository: productRepository}
}
