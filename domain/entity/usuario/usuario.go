package entity

type Usuario struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Idade    string      `json:"idade"`
	Endereco []*Endereco `json:"endereco"`
	Cpf      string      `json:"cpf"`
	senha    string      `json:"senha"`
	Telefone string      `json:"Telefone"`
}

func NewUsuario(name, email, telefone, senha string) *Usuario {

	return &Usuario{Name: name,
		senha:    senha,
		Email:    email,
		Telefone: telefone,
	}

}

func (u *Usuario) SetEndereco(endereco ...*Endereco) {
	u.Endereco = endereco
}
