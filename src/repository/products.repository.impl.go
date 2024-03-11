package repository

import (
	"errors"
	"github.com/luthfiybk/go-with-gin/src/model"
	"github.com/luthfiybk/go-with-gin/src/data/request"
	"github.com/luthfiybk/go-with-gin/src/helper"

	"gorm.io/gorm"
)

type ProductsRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductsRepositoryImpl (Db *gorm.DB) ProductsRepository {
	return &ProductsRepositoryImpl{Db: Db}
}

func (p ProductsRepositoryImpl) Update(products model.Products) {
	var updateProduct = request.UpdateProductsRequest{
		Id: products.Id,
		Product_Name: products.Product_Name,
		Description: products.Description,
	}
	result := p.Db.Model(&products).Updates(updateProduct)
	helper.ErrorPanic(result.Error)
}

func (p ProductsRepositoryImpl) Save(products model.Products) {
	result := p.Db.Create(&products)
	helper.ErrorPanic(result.Error)
}

func (p ProductsRepositoryImpl) Delete(productsId int) {
	var products model.Products
	result := p.Db.Where("id = ?", productsId).Delete(&products)
	helper.ErrorPanic(result.Error)
}

func (p ProductsRepositoryImpl) FindById(productsId int) (model.Products, error) {
	var product model.Products
	result := p.Db.Find(&product, productsId)
	if result != nil {
		return product, nil
	} else {
		return product, errors.New("tag is not found")
	}
}

func (p ProductsRepositoryImpl) FindAll() []model.Products {
	var products []model.Products
	results := p.Db.Find(&products)
	helper.ErrorPanic(results.Error)
	return products
}