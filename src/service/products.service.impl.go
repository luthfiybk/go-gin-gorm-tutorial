package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/luthfiybk/go-with-gin/src/data/request"
	"github.com/luthfiybk/go-with-gin/src/data/response"
	"github.com/luthfiybk/go-with-gin/src/helper"
	"github.com/luthfiybk/go-with-gin/src/model"
	"github.com/luthfiybk/go-with-gin/src/repository"
)

type ProductsServiceImpl struct {
	ProductRepository repository.ProductsRepository
	Validate *validator.Validate
}

func NewProductsServiceImpl (productRepository repository.ProductsRepository, validate *validator.Validate) ProductsService {
	return &ProductsServiceImpl{
		ProductRepository: productRepository,
		Validate: validate,
	}
}

func (p ProductsServiceImpl) Create(product request.CreateProductsRequest) {
	err := p.Validate.Struct(product)
	helper.ErrorPanic(err)
	productModel := model.Products{
		Product_Name: product.Product_Name,
		Description: product.Description,
	}
	p.ProductRepository.Save(productModel)
}

func (p ProductsServiceImpl) Update(product request.UpdateProductsRequest) {
	productData, err := p.ProductRepository.FindById(product.Id)
	helper.ErrorPanic(err)
	productData.Product_Name = product.Product_Name
	p.ProductRepository.Update(productData)
}

func (p ProductsServiceImpl) Delete(productId int) {
	p.ProductRepository.Delete(productId)
}

func (p ProductsServiceImpl) FindById(productId int) response.ProductsResponse {
	productData, err := p.ProductRepository.FindById(productId)
	helper.ErrorPanic(err)

	productResponse := response.ProductsResponse{
		Id: productData.Id,
		Product_Name: productData.Product_Name,
		Description: productData.Description,
	}

	return productResponse
}

func (p ProductsServiceImpl) FindAll() []response.ProductsResponse {
	productsData := p.ProductRepository.FindAll()
	var productsResponse []response.ProductsResponse
	for _, product := range productsData {
		productResponse := response.ProductsResponse{
			Id: product.Id,
			Product_Name: product.Product_Name,
			Description: product.Description,
		}
		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse
}