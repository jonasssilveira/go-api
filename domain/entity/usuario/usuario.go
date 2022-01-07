package entity

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Id       string  `json:"id";gorm:"type:uuid;primaryKey;id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Idade    string     `json:"idade"`
	Endereco []Endereco `json:"endereco" gorm:"foreignKey:ID;references:ID"`
	Cpf      string     `json:"cpf"`
	senha    string     `json:"senha"`
	Telefone string     `json:"Telefone"`
}

func (u *Usuario) SetEndereco(endereco ...Endereco) {
	u.Endereco = endereco
}
