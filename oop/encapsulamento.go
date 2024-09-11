package oop

type Usuario struct {
	nome, email, cpf string
}

func (u *Usuario) Nome() string {
	return u.nome
}

func (u *Usuario) SetNome(nome string) {
	u.nome = nome
}
