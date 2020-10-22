package handler

import (
	"net/http"

	"github.com/yaminmhd/go-hardware-store/constant"
	"github.com/yaminmhd/go-hardware-store/contract"
	"github.com/yaminmhd/go-hardware-store/log"
)

type Product struct {
	service ProductService
}

func (handler *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	products, err := handler.service.GetProducts(ctx)
	if err != nil {
		log.Log.Errorf(r, "[Products] Could not fetch products, %s", err)
		status := contract.ErrorObjects[err.Error()].Status
		contract.ErrorResponse(w, []string{err.Error()}, status)
		return
	}
	successStatus := constant.SuccessOK
	contract.SuccessfulResponse(w, products, successStatus)
}

func NewProductHandler(service ProductService) Product {
	return Product{
		service: service,
	}
}
