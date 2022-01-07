package interfaces

import (
	"e-commerce/domain/entity/produto"
)

type ProductRepository interface {
	ListAllProducts() []produto.Product
	ListAllProductsByCategory() []produto.Product
	GetProductById(id string) *produto.Product
	GetProductByName(name string) (*produto.Product, error)
	GetProductByNameOrCategoryName(name string) *produto.Product
	SaveProduct(product *produto.Product) *produto.Product
	UpdateProduct(product *produto.Product) *produto.Product
	GetProductByVenda(id string) (*produto.Product,error)
	DeleteProduct(id string) error
}
