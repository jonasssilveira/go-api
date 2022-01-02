package produto

import (
	"errors"
	"github.com/satori/go.uuid"
)

type Product struct {
	Id        string
	Name      string
	Quant     int
	Price     float64
	Descricao string
	Categoria
	Fabricante
}

func NewProduct(name, descricao string, quant int, price float64, categoria Categoria, fabricante Fabricante) *Product {
	return &Product{
		Id:         uuid.NewV4().String(),
		Name:       name,
		Descricao:  descricao,
		Price:      price,
		Quant:      quant,
		Categoria:  categoria,
		Fabricante: fabricante,
	}
}

func (p *Product) Sell() error {
	if p.Quant > 0 {
		p.Quant--
		return nil
	}
	return errors.New("Quantidade insuficiente para ser vendida")
}

func (p *Product) Buy(i int) {
	p.Quant += i
}
