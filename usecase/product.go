package usecase

import (
	"e-commerce/domain/entity/produto"
	"e-commerce/usecase/interfaces"
	"errors"
)

type Product struct {
	Repository interfaces.ProductRepository
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
	prod := p.Repository.GetProductByVenda(id)
	if prod != nil {
		return p.Repository.DeleteProduct(id)
	}
	return errors.New("Não foi possivel deletar pois existem vendas deste produto")
}
func (p *Product) SaveProduct(product *produto.Product) (*produto.Product, error) {
	prod := p.Repository.GetProductByName(product.Name)
	if prod != nil {
		return nil, errors.New("Não é possivel salvar, pois o produto já existe!")
	}
	return p.Repository.SaveProduct(product), nil
}

func (p *Product) Search(product *produto.Product) (*produto.Product, error) {
	prod := p.Repository.GetProductByNameOrCategoryName(product.Name)
	if prod != nil {
		return nil, errors.New("Não foi encontrado nenhum produto com estas caracteristicas")
	}
	return prod, nil
}
