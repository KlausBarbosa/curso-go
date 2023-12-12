package main

import (
	"fmt"
	"github.com/curso-go/pkg1/matematica"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Printf("Resultado: %v ", s)
}
