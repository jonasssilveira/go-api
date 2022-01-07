package usecase

import (
	"e-commerce/domain/entity/produto"
	"e-commerce/usecase/mock"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var categoria = produto.Categoria{Id: "123", Name: "Filmadora"}
var fabricante = produto.Fabricante{Id: "123", Name: "TecnoMania"}

var product = produto.NewProduct("Filmadora",
	"TekPix, a câmera mais vendida do Brasil",
	10, 1000, categoria, fabricante)

func TestIfDeleteWithSuccess(t *testing.T) {
	//arrange
	mockRepo := new(mock.MockRepository)
	id := uuid.NewV4().String()
	mockRepo.On("DeleteProduct").Return(nil)
	mockRepo.On("GetProductByVenda", id).Return(product, nil)

	//act
	prod := NewProduct(mockRepo)
	err := prod.DeleteProduct(id)

	//assert
	assert.Nil(t, err)
}
func TestIfDeleteWithError(t *testing.T) {
	//arrange
	mockRepo := new(mock.MockRepository)
	id := uuid.NewV4().String()
	er := errors.New("Não foi possivel deletar pois existem vendas deste produto")
	mockRepo.On("GetProductByVenda", id).Return(&produto.Product{}, er)

	//act
	prod := NewProduct(mockRepo)
	err := prod.DeleteProduct(id)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, er, err)
}

func TestIfSaveWithSuccess(t *testing.T) {
	//arrange
	mockRepo := new(mock.MockRepository)
	mockRepo.On("GetProductByName", "Filmadora").Return(product, nil)
	mockRepo.On("SaveProduct", product).Return(product)

	//act
	prod := NewProduct(mockRepo)
	saveProduct, err := prod.SaveProduct(product)
	//assert
	assert.NotNil(t, saveProduct)
	assert.Nil(t, err)
}
func TestIfNotSave(t *testing.T) {
	//arrange
	mockRepo := new(mock.MockRepository)

	er := errors.New("Não é possivel salvar, pois o produto já existe!")
	mockRepo.On("GetProductByName", "Filmadora").Return(product, er)

	//act
	prod := NewProduct(mockRepo)
	saveProduct, err := prod.SaveProduct(product)
	//assert
	assert.Nil(t, saveProduct)
	assert.NotNil(t, err)
	assert.Equal(t, er, err)
}
func TestIfSearchReturnSuccessfull(t *testing.T) {
	//arrange
	mockRepo := new(mock.MockRepository)

	mockRepo.On("GetProductByNameOrCategoryName", "Filmadora").Return(product)

	//act
	prod := NewProduct(mockRepo)
	productSeached, err := prod.Search(product.Name)

	//assert
	assert.Equal(t, product.Id, productSeached.Id)
	assert.Nil(t, err)
}
func TestIfSearchReturnNotSuccessfull(t *testing.T) {
	//arrange
	mockRepo := new(mock.MockRepository)

	er := errors.New("Não foi encontrado nenhum produto com estas caracteristicas")
	mockRepo.On("GetProductByNameOrCategoryName", "Filmadora").Return(&produto.Product{})

	//act
	prod := NewProduct(mockRepo)
	productSeached, err := prod.Search(product.Name)
	//assert
	assert.Nil(t, productSeached)
	assert.NotNil(t, err)
	assert.Equal(t, er, err)
}
