package payment

import (
	entity2 "e-commerce/domain/entity/produto"
	"e-commerce/domain/entity/usuario"
	uuid "github.com/satori/go.uuid"
)

type Carrinho struct {
	Id string
	Usuario
	entity.Endereco
	Product  []*entity2.Product
	Amount   float64
	Desconto string
}

func NewCarrinho(usuario *Usuario, desconto string, products ...*entity2.Product) *Carrinho {
	carrinho := &Carrinho{
		Id:       uuid.NewV4().String(),
		Usuario:  *usuario,
		Endereco: usuario.Endereco,
		Product:  products,
		Desconto: desconto,
	}

	var amount float64
	for _, product := range products {
		amount += product.Price
	}
	carrinho.Amount = amount
	return carrinho
}

func (c *Carrinho) ExecuteSell() error {
	for _, product := range c.Product {
		err := product.Sell()
		if err != nil {
			return err
		}
	}
	return nil
}
