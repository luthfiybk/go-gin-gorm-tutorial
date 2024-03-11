package service

import (
	"github.com/luthfiybk/go-with-gin/src/data/request"
	"github.com/luthfiybk/go-with-gin/src/data/response"
)

type ProductsService interface {
	Create(products request.CreateProductsRequest)
	Update(products request.UpdateProductsRequest)
	Delete(productsId int)
	FindById(productsId int) response.ProductsResponse
	FindAll() []response.ProductsResponse
}