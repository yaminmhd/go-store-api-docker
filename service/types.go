package service

import (
	"context"

	"github.com/yaminmhd/go-hardware-store/model"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]*model.Product, error)
}
