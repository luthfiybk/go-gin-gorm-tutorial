package router

import (
	"github.com/luthfiybk/go-with-gin/src/controller"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(productsController *controller.ProductsController) *gin.Engine {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	productsRouter := baseRouter.Group("/products")
	productsRouter.GET("", productsController.FindAll)
	productsRouter.GET("/:productId", productsController.FindById)
	productsRouter.POST("", productsController.Create)
	productsRouter.PATCH("/:productId", productsController.Update)
	productsRouter.DELETE("/:productId", productsController.Delete)

	return router
}