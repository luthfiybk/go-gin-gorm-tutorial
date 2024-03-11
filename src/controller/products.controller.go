package controller

import (
	"github.com/luthfiybk/go-with-gin/src/data/request"
	"github.com/luthfiybk/go-with-gin/src/data/response"
	"github.com/luthfiybk/go-with-gin/src/helper"
	"github.com/luthfiybk/go-with-gin/src/service"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ProductsController struct {
	productsService service.ProductsService
}

func NewProductsController(service service.ProductsService) *ProductsController {
	return &ProductsController{productsService: service}
}

func (controller *ProductsController) Create(ctx *gin.Context) {
	log.Info().Msg("Create product")
	createProductsRequest := request.CreateProductsRequest{}
	err := ctx.ShouldBindJSON(&createProductsRequest)
	helper.ErrorPanic(err)

	controller.productsService.Create(createProductsRequest)
	webResponse := response.Response{
		Code: http.StatusOK,
		Data: nil,
		Status: "Success",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductsController) Update(ctx *gin.Context) {
	log.Info().Msg("Update Product")
	updateProductRequest := request.UpdateProductsRequest{}
	err := ctx.ShouldBindJSON(&updateProductRequest)
	helper.ErrorPanic(err)

	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)
	updateProductRequest.Id = id

	controller.productsService.Update(updateProductRequest)
	webResponse := response.Response{
		Code: http.StatusOK,
		Data: nil,
		Status: "Success",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func  (controller *ProductsController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete Product")
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)
	controller.productsService.Delete(id)

	webResponse := response.Response{
		Code: http.StatusOK,
		Data: nil,
		Status: "Success",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductsController) FindById(ctx *gin.Context) {
	log.Info().Msg("Find Product By ID")
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)
	productResponse := controller.productsService.FindById(id)

	webResponse := response.Response{
		Code: http.StatusOK,
		Data: productResponse,
		Status: "Success",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductsController) FindAll(ctx *gin.Context) {
	log.Info().Msg("Find All Product")
	productResponse := controller.productsService.FindAll()
	
	webResponse := response.Response{
		Code: http.StatusOK,
		Data: productResponse,
		Status: "Success",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}