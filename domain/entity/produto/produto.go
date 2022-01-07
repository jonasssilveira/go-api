package produto

import (
	"errors"
	"github.com/satori/go.uuid"
	"time"
)

type Product struct {
	Id        string    `json:"id"gorm:"type:uuid;primaryKey;id"`
	Name      string    `json:"name"`
	Quant     int       `json:"quant"`
	Price     float64   `json:"price"`
	Descricao string    `json:"descricao"`
	Categoria Categoria `json:"categoria"gorm:"foreignKey:ID;references:ID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewProduct(name, descricao string, quant int, price float64, categoria Categoria, fabricante Fabricante) *Product {
	return &Product{
		Id:        uuid.NewV4().String(),
		Name:      name,
		Descricao: descricao,
		Price:     price,
		Quant:     quant,
		Categoria: categoria,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ProductNew() *Product {
	return &Product{
		Id:        uuid.NewV4().String(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
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
