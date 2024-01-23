package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch) //não irá rodar em background pois o programa morreria (não há nada segurando o processo)

}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //sinaliza que o canal fechou apos o processamento, impedindo deadlock
}
