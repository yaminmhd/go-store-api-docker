package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/yaminmhd/go-hardware-store/model"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (mock *ProductRepositoryMock) GetAllProducts(ctx context.Context) ([]*model.Product, error){
	args := mock.Called()
	return args.Get(0).([]*model.Product), args.Error(1)
}
