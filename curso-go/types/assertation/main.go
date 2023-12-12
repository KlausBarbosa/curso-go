package main

import "fmt"

func main() {
	var minhaVar interface{} = "Jovirone Fonseco"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res eh %v e o resultado de ok eh %v", res, ok)

}
