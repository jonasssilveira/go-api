package produto

import (
	"errors"
	"github.com/satori/go.uuid"
)

type Product struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Quant      int     `json:"quant"`
	Price      float64 `json:"price''"`
	Descricao  string  `json:"descricao"`
	Categoria  `json:"categoria"`
	Fabricante `json:"fabricante"`
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
