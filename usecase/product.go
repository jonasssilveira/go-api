package usecase

import (
	"e-commerce/domain/entity/produto"
	"e-commerce/usecase/interfaces"
	"errors"
)

type Product struct {
	Repository interfaces.ProductRepository
}

func NewProduct(repository interfaces.ProductRepository)*Product{
	return &Product{
		Repository:repository,
	}
}

func (p *Product) ListProducts() []produto.Product {
	return p.Repository.ListAllProducts()
}
func (p *Product) ListProductsByCategory() []produto.Product {
	return p.Repository.ListAllProductsByCategory()
}
func (p *Product) GetByID(id string) *produto.Product {
	return p.Repository.GetProductById(id)
}
func (p *Product) UpdateProduct(product *produto.Product) *produto.Product {
	return p.Repository.UpdateProduct(product)
}
func (p *Product) DeleteProduct(id string) error {
	prod, error := p.Repository.GetProductByVenda(id)
	if error != nil {
		return errors.New("Não foi possivel deletar pois existem vendas deste produto")
	}
	return p.Repository.DeleteProduct(prod.Id)
}

func (p *Product) SaveProduct(product *produto.Product) (*produto.Product, error) {
	error := p.Repository.GetProductByName(product.Name)
	if error != nil {
		return nil, errors.New("Não é possivel salvar, pois o produto já existe!")
	}
	return p.Repository.SaveProduct(product), nil
}

func (p *Product) Search(name string) (*produto.Product, error) {
	prod := p.Repository.GetProductByNameOrCategoryName(name)
	if prod.Id == "" {
		return nil, errors.New("Não foi encontrado nenhum produto com estas caracteristicas")
	}
	return prod, nil
}
