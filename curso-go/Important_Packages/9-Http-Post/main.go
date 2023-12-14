package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Carlos"}`)) //necessario jogar o slic de bytes em um reader/bufferizar
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil) //ira pegar os dados, informamos de onde vem os dados e pra onde eles serao copiados (resp.body)
}
