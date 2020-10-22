package handler

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/yaminmhd/go-hardware-store/contract"
)

type ProductServiceMock struct {
	mock.Mock
}

func (mock *ProductServiceMock) GetProducts(ctx context.Context) (contract.GetProducts, error) {
	args := mock.Called()
	return args.Get(0).(contract.GetProducts), args.Error(1)
}
