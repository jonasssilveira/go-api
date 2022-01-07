package controllers

import (
	"e-commerce/adapter/database/repository"
	"e-commerce/domain/entity/produto"
	"e-commerce/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	product *usecase.Product
	db      *gorm.DB
}

func NewProductController() *ProductController {
	db := repository.StartDB()
	productRepo := repository.NewProductRepository(db)
	product := usecase.NewProduct(productRepo)
	return &ProductController{
		product: product,
	}
}

func (cont *ProductController) ListAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		products := cont.product.ListProducts()
		c.JSON(200, products)
	}
}

func (cont *ProductController) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product := cont.product.GetByID(id)
		c.JSON(200, product)
	}
}
func (cont *ProductController) SaveProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod := produto.ProductNew()
		err := c.ShouldBind(prod)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
		}
		saveProduct, err := cont.product.SaveProduct(prod)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
		}
		c.JSON(200, saveProduct)

	}
}

func (cont *ProductController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := cont.product.DeleteProduct(id)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Não foi possivel salvar" + err.Error(),
			})
		}
		c.JSON(200, gin.H{
			"success": id + " foi deletado com sucesso!",
		})
	}
}
