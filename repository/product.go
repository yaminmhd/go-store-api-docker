package repository

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/yaminmhd/go-hardware-store/constant"
	"github.com/yaminmhd/go-hardware-store/log"
	"github.com/yaminmhd/go-hardware-store/model"
)

type Product struct {
	db *gorm.DB
}


func(repo Product) GetAllProducts(ctx context.Context) ([]*model.Product, error){
	var products []*model.Product

	err := repo.db.Debug().Model(&model.Product{}).Find(&products).Error
	if err != nil{
		log.Log.Error("[ProductRepository] error getting products from DB", err)
		return nil, errors.New(constant.ErrorInternalServerError)
	}
	return products, err
}

func NewProductRepository(db *gorm.DB) Product{
	return Product{
		db: db,
	}
}