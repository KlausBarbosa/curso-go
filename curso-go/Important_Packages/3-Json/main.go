package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero_conta"`
	Saldo  int `json:"saldo_valor"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	//Marshal no caso ed transformar/serializar em Json

	if err != nil {
		panic(err)
	}
	println(string(res))
	//importante utilizar o cast pra string, pois o json sempre retornara em bytes.

	err = json.NewEncoder(os.Stdout).Encode(conta)
	//encoder recebe um valor e faz um 'encoding' gravando em outro lugar (arquivo, tela do terminal, etc exemplo do os.stdout que eh a saida do terminal)

	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	//pega o valor de jsonPuro e joga no endereco de memoria onde contaX esta armazenado, alterando seu valor. Unmarshal tambem retorna um erro, caso houver.
	if err != nil {
		panic(err)
	}
	println(contaX.Saldo)

}
