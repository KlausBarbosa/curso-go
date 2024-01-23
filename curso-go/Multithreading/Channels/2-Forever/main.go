package main

// Thread 1
func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever //espera o canal ficar cheio para que este seja esvaziado, resulta em um deadlock caso nao tenha valor
}
