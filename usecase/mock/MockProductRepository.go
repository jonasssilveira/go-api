package mock

import (
	"e-commerce/domain/entity/produto"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}


func (mock *MockRepository) ListAllProducts() []produto.Product {
	args := mock.Called()
	result := args.Get(0)
	return result.([]produto.Product)
}
func (mock *MockRepository) ListAllProductsByCategory() []produto.Product {
	args := mock.Called()
	result := args.Get(0)
	return result.([]produto.Product)
}
func (mock *MockRepository) GetProductById(id string) *produto.Product {
	args := mock.Called()
	result := args.Get(0)
	return result.(*produto.Product)
}
func (mock *MockRepository) GetProductByName(name string) error {
	args := mock.Called(name)
	return  args.Error(0)
}
func (mock *MockRepository) GetProductByNameOrCategoryName(name string) *produto.Product {
	args := mock.Called(name)
	result := args.Get(0)
	return result.(*produto.Product)
}
func (mock *MockRepository) SaveProduct(product *produto.Product) *produto.Product {
	args := mock.Called(product)
	result := args.Get(0)
	return result.(*produto.Product)}
func (mock *MockRepository) UpdateProduct(product *produto.Product) *produto.Product {
	args := mock.Called()
	result := args.Get(0)
	return result.(*produto.Product)
}
func (mock *MockRepository) GetProductByVenda(id string) (*produto.Product, error) {
	args := mock.Called(id)
	result := args.Get(0)
	return result.(*produto.Product), args.Error(1)
}
func (mock *MockRepository) DeleteProduct(id string) error {
	args := mock.Called()
	return args.Error(0)
}
