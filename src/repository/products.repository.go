package repository

import (
	"github.com/luthfiybk/go-with-gin/src/model"
)

type ProductsRepository interface {
	Save(products model.Products)
	Update(products model.Products)
	Delete(productsId int)
	FindById(productsId int) (tags model.Products, err error)
	FindAll() []model.Products
}