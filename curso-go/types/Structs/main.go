package main

import "fmt"

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	wesley.Ativo = false
	//wesley.Address.Cidade = "São Vicente"
	//wesley.Address.Logradouro = "Marques do Seu Vicente"
	//wesley.Address.Numero = 1857
	//wesley.Desativar()
	//Desativacao(wesley)

	//fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)

	minhaEmpresa := Empresa{"GUniverse"}
	Desativacao(minhaEmpresa)
}
