package handler

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"github.com/yaminmhd/go-hardware-store/config"
	"github.com/yaminmhd/go-hardware-store/contract"
	"github.com/yaminmhd/go-hardware-store/log"
)

type ProductHandlerProductSuite struct {
	ctx context.Context
	suite.Suite
	handler     Product
	mockService *ProductServiceMock
}

func (suite *ProductHandlerProductSuite) SetupSuite() {
	config.Load()
	log.SetupLogger()
}

func (suite *ProductHandlerProductSuite) SetupTest() {
	suite.ctx = context.TODO()
	suite.mockService = &ProductServiceMock{}
	suite.handler = NewProductHandler(suite.mockService)
}

func (suite *ProductHandlerProductSuite) TestGetHandlerShouldReturnSuccess() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/products", suite.handler.GetProducts).Methods(http.MethodGet)
	expectedResponse := `{"success":true,"data":{"products":[{"id":123,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","name":"Powersaw","price":54.5,"quantity":5,"state":"available"},{"id":456,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","name":"Screwdriver","price":11.99,"quantity":3,"state":"available"}],"total_count":2}}`

	serviceResponse := contract.GetProducts{
		Products: []contract.ProductSummary{
			{ID: 123, Name: "Powersaw", Price: 54.50, Quantity: 5, State: "available"},
			{ID: 456, Name: "Screwdriver", Price: 11.99, Quantity: 3, State: "available"},
		},
		TotalCount: 2,
	}
	suite.mockService.On("GetProducts").Return(serviceResponse, nil)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/v1/products", nil)
	request = request.WithContext(suite.ctx)

	router.ServeHTTP(recorder, request)
	responseBytes, _ := ioutil.ReadAll(recorder.Body)

	suite.Equal(http.StatusOK, recorder.Code)
	suite.Equal(expectedResponse, string(responseBytes))
	suite.mockService.AssertExpectations(suite.T())
}

func TestProductHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerProductSuite))
}
