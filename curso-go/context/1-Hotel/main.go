package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	//informa que o contexto foi iniciado (em branco) e depois que iniciado, ira rodar na thread inicial

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	//Temos withCancel, WithDeadline(), WithTimeout(countdown), WithValue(quando o valor eh passado)

	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	//	select eh similar ao case, porem, ele espera uma acao para poder ser executado
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
