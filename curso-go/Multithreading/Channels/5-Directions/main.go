package main

import "fmt"

func recebe(nome string, hello chan<- string) { //a <- indica que o canal ira apenas receber info
	hello <- nome
}

func ler(data <-chan string) {
	fmt.Println(<-data) //esvazia o canal
}

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
