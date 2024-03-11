package main

import (
	"github.com/luthfiybk/go-with-gin/src/config"
	"github.com/luthfiybk/go-with-gin/src/controller"
	"github.com/luthfiybk/go-with-gin/src/helper"
	"github.com/luthfiybk/go-with-gin/src/model"
	"github.com/luthfiybk/go-with-gin/src/repository"
	"github.com/luthfiybk/go-with-gin/src/router"
	"github.com/luthfiybk/go-with-gin/src/service"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main(){
	log.Info().Msg("Starting application")
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("products").AutoMigrate(&model.Products{})

	// Repository
	productsRespository := repository.NewProductsRepositoryImpl(db)

	// Service
	productsService := service.NewProductsServiceImpl(productsRespository, validate)

	// Controller
	productsController := controller.NewProductsController(productsService)

	// Router
	routes := router.NewRouter(productsController)

	server := &http.Server{
		Addr: ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
