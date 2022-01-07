package entity

import (
	"gorm.io/gorm"
)

type Endereco struct {
	gorm.Model
	ID          string  `gorm:"type:uuid;primaryKey;id"`
	Logradouro  string  `json:"logradouro"`
	Numero      int     `json:"numero"`
	Complemento string  `json:"complemento"`
	Bairro      string  `json:"bairro"`
	Cidade      string  `json:"cidade"`
	CEP         int     `json:"cep"`
	UF          string  `json:"uf"`
}

func NewEndereco(complemento, bairro, cidade, logadouro, uf string, numero, cep int) *Endereco {
	return &Endereco{
		CEP:         cep,
		Numero:      numero,
		Cidade:      cidade,
		Complemento: complemento,
		Bairro:      bairro,
		Logradouro:  logadouro,
		UF:          uf,
	}
}
