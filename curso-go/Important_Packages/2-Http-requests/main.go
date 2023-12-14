package main

import (
	"io"
	"net/http"
)

func main() {

	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()
	//fara com que o Close ocorra somente depois que toda a func main rodar.

	res, err := io.ReadAll(request.Body)
	//read all pega o que ta dentro do Body e faz a leitura do stream em processamento
	if err != nil {
		panic(err)
	}
	println(string(res))
	//	O defer rodara aqui
}
