package payment

import (
	entity2 "e-commerce/domain/entity"
	"e-commerce/domain/entity/produto"
	"e-commerce/domain/entity/usuario"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfCarrinhoIsNotEmpty(t *testing.T) {
	var product1 = produto.NewProduct("Filmadora",
		"TekPix, a câmera mais vendida do Brasil",
		66, 1000, entity2.categoria, entity2.fabricante)
	var product2 = produto.NewProduct("Filmadora",
		"TekPix, a câmera mais vendida do Brasil",
		66, 1000, entity2.categoria, entity2.fabricante)
	var endereco = entity.NewEndereco("Bloco 7 ap 202", "Anchieta", "Rio de Janeiro",
		"AV", "RJ", 1073, 21645521)

	var usuario = NewUsuario("Jonas", "jonasdasilvasilveira@gmail.com", "21965224950", "123")
	usuario.SetEndereco(endereco)
	var carrinho = NewCarrinho(usuario, "", product1, product2)

	assert.Equal(t, product1.Id, carrinho.Product[0].Id)
	assert.Equal(t, product2.Id, carrinho.Product[1].Id)
	assert.Equal(t, usuario.Id, carrinho.Usuario.Id)
	assert.Equal(t, 2000., carrinho.Amount)
}
func TestIfCarrinhoIsSelling(t *testing.T) {
	var product1 = produto.NewProduct("Filmadora",
		"TekPix, a câmera mais vendida do Brasil",
		5, 1000, entity2.categoria, entity2.fabricante)
	var endereco = entity.NewEndereco("Bloco 7 ap 202", "Anchieta", "Rio de Janeiro",
		"AV", "RJ", 1073, 21645521)

	var usuario = NewUsuario("Jonas", "jonasdasilvasilveira@gmail.com", "21965224950", "123")
	usuario.SetEndereco(endereco)
	var carrinho = NewCarrinho(usuario, "", product1, product1, product1, product1, product1)
	assert.Nil(t, carrinho.ExecuteSell())
	assert.Equal(t, product1.Id, carrinho.Product[0].Id)
	assert.Equal(t, 0, product1.Quant)
	assert.Equal(t, usuario.Id, carrinho.Usuario.Id)
	assert.Equal(t, 5000., carrinho.Amount)
}

func TestIfCarrinhoIsNotSelling(t *testing.T) {
	var product1 = produto.NewProduct("Filmadora",
		"TekPix, a câmera mais vendida do Brasil",
		2, 1000, entity2.categoria, entity2.fabricante)
	var endereco = entity.NewEndereco("Bloco 7 ap 202", "Anchieta", "Rio de Janeiro",
		"AV", "RJ", 1073, 21645521)

	var usuario = NewUsuario("Jonas", "jonasdasilvasilveira@gmail.com", "21965224950", "123")
	usuario.SetEndereco(endereco)
	var carrinho = NewCarrinho(usuario, "", product1, product1, product1, product1, product1)
	assert.NotNil(t, carrinho.ExecuteSell())
	assert.Equal(t, product1.Id, carrinho.Product[0].Id)
	assert.Equal(t, usuario.Id, carrinho.Usuario.Id)
	assert.Equal(t, 5000., carrinho.Amount)
}
