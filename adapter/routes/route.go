package routes

import (
	"e-commerce/adapter/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	productController := controllers.NewProductController()
	main := router.Group("api/v1")
	{
		products := main.Group("product")
		{
			products.GET("/", productController.ListAllProducts())
			products.GET("/:id", productController.GetProductById())
			products.POST("/", productController.SaveProduct())
			products.DELETE("/:id", productController.Delete())
		}
	}
	return router
}
