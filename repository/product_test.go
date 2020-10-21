package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"github.com/yaminmhd/go-hardware-store/appcontext"
	"github.com/yaminmhd/go-hardware-store/config"
	"github.com/yaminmhd/go-hardware-store/log"
	"github.com/yaminmhd/go-hardware-store/model"
	"testing"
	"time"
)

type ProductRepositoryTestSuite struct {
	suite.Suite
	repository Product
	db         *gorm.DB
	ctx        context.Context
}

func (suite *ProductRepositoryTestSuite) SetupSuite() {
	config.Load()
	log.SetupLogger()
}

func (suite *ProductRepositoryTestSuite) SetupTest() {
	appcontext.Initiate()
	suite.db = appcontext.GetDB()
	suite.db.Debug().CreateTable(&model.Product{})
	suite.repository = NewProductRepository(suite.db)
}

func (suite *ProductRepositoryTestSuite) TearDownTest() {
	suite.db.Debug().DropTableIfExists(&model.Product{})
	_ = appcontext.GetDB().Close()
}

func (suite *ProductRepositoryTestSuite) insertTestDataProduct(productData []*model.Product) ([]*model.Product, error) {
	transaction := suite.db.Begin()
	var product []*model.Product
	for _,product := range productData {
		err := transaction.Debug().Model(&model.Product{}).Create(&product).Error
		if err != nil {
			transaction.Rollback()
			return nil, err
		}
	}
	return product, transaction.Commit().Error
}

func (suite *ProductRepositoryTestSuite) TestGetAllProductsShouldReturnProductsSuccessfully() {
	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	product1 := &model.Product{
		ID:        123,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      "Saw",
		Price:     15.50,
		Quantity:  5,
		State:     "available",
	}
	product2 := &model.Product{
		ID:        456,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      "Hammer",
		Price:     11.50,
		Quantity:  3,
		State:     "available",
	}
	products := []*model.Product{
		product1,product2,
	}
	_, err := suite.insertTestDataProduct(products)
	actualProducts, err := suite.repository.GetAllProducts(suite.ctx)

	suite.Equal(products, actualProducts)
	suite.Equal(2, len(actualProducts))
	suite.NoError(err)
}

func (suite *ProductRepositoryTestSuite) TestGetAllProductsShouldReturnEmptyProductsSuccessfully() {
	actualProducts, err := suite.repository.GetAllProducts(suite.ctx)
	suite.Equal([]*model.Product{}, actualProducts)
	suite.Equal(0, len(actualProducts))
	suite.NoError(err)
}

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}
