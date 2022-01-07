package payment

import (
	entity2 "e-commerce/domain/entity/produto"
	"e-commerce/domain/entity/usuario"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Carrinho struct {
	gorm.Model
	Id       string             `json:"id"gorm:"type:uuid;primaryKey;id"`
	Usuario  entity.Usuario     `json:"usuario"gorm:"foreignKey:Id;references:Id"`
	Endereco entity.Endereco    `json:"endereco"gorm:"foreignKey:ID;references:ID"`
	Product  []*entity2.Product `json:"product"gorm:"foreignKey:Id;references:Id"`
	Amount   float64            `json:"amount"`
	Desconto string             `json:"desconto"`
}

func NewCarrinho(usuario *entity.Usuario, desconto string, products ...*entity2.Product) *Carrinho {
	carrinho := &Carrinho{
		Id:       uuid.NewV4().String(),
		Usuario:  *usuario,
		Endereco: usuario.Endereco[0],
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
