package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) //Vazio

	// Thread 2
	go func() {
		channel <- "Hello World!" //Cheio
	}()

	// Thread 1
	msg := <-channel //Vazio
	fmt.Println(msg)
}
