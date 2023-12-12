package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()                                 //criacao de um contexto vazio
	ctx, cancel := context.WithTimeout(ctx, 1*time.Microsecond) //se passar um segundo do processamento, o contexto sera cancelado.
	//ctx, cancel := context.WithCancel(ctx)  //tambem possivel utilizar esta funcao dos contextos, o acionamento sera via defer/chamada direta atraves de um if, etc.
	defer cancel()                                                               //como boa pratica, colocamos o cancel de toda forma, pra caso de algum problema no cancelamento por contexto.
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil) //Iremos passar o contexto configurado para esta request, e so ira ser continuada ed acordo com a regua passada anteriormente no context.
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

}
