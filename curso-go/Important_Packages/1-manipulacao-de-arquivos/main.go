package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//tamanho, err := f.WriteString("Hello, World!")
	tamanho, err := f.Write([]byte("Dados escritos no arquivo"))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", tamanho)
	f.Close()

	//LEITURA DE ARQUIVO
	/*	arquivo, err := os.ReadFile("arquivo.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println("\nLendo dados no arquivo", string(arquivo))*/

	//	leitura de pouco em pouco abrindo o arquivo
	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 3) //o buffer ira ler de 10 em 10 bytes

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
