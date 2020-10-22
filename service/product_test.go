package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/yaminmhd/go-hardware-store/config"
	"github.com/yaminmhd/go-hardware-store/contract"
	"github.com/yaminmhd/go-hardware-store/log"
	"github.com/yaminmhd/go-hardware-store/model"
)

type ProductServiceTestSuite struct {
	suite.Suite
	ctx               context.Context
	productRepository *ProductRepositoryMock
	service           ProductService
}

func (suite *ProductServiceTestSuite) SetupSuite() {
	config.Load()
	log.SetupLogger()
}

func (suite *ProductServiceTestSuite) SetupTest() {
	suite.ctx = context.TODO()
	suite.productRepository = &ProductRepositoryMock{}
	suite.service = NewProductService(suite.productRepository)
}

func (suite *ProductServiceTestSuite) TestGetAllProductsShouldReturnAllProductsSuccessfully() {

	product1 := &model.Product{
		ID:       123,
		Name:     "Saw",
		Price:    10.00,
		Quantity: 5,
		State:    "available",
	}

	product2 := &model.Product{
		ID:       456,
		Name:     "Screwdriver",
		Price:    5.90,
		Quantity: 2,
		State:    "available",
	}

	productsFromDb := []*model.Product{
		product1, product2,
	}

	expectedProducts := contract.GetProducts{
		Products: []contract.ProductSummary{
			{
				ID:       123,
				Name:     "Saw",
				Price:    10.00,
				Quantity: 5,
				State:    "available",
			},
			{
				ID:       456,
				Name:     "Screwdriver",
				Price:    5.90,
				Quantity: 2,
				State:    "available",
			},
		},
		TotalCount: 2,
	}
	suite.productRepository.On("GetAllProducts").Return(productsFromDb, nil)

	actualProducts, err := suite.service.GetProducts(suite.ctx)
	suite.Equal(expectedProducts, actualProducts)
	suite.NoError(err)
	suite.productRepository.AssertExpectations(suite.T())
}

func TestProductServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}
