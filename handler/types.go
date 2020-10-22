package handler

import (
	"context"

	"github.com/yaminmhd/go-hardware-store/contract"
)

type ProductService interface {
	GetProducts(ctx context.Context) (contract.GetProducts, error)
}
