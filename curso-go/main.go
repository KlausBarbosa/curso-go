package main

import (
	"errors"
	"fmt"
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 10

	func() int {
		return sumVar(1, 3, 4, 5) * 2
	}()

	//fmt.Println(meuArray[len(meuArray)-1])

	//for i, v := range meuArray {
	//	fmt.Printf("O valor do índice %d é %d\n", i, v)
	//}

	//s := []int{2, 4, 6, 8, 10}
	//fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	//fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	//fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	//fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])
	//
	//s = append(s, 110)

	salarios := map[string]int{"Wesley": 1000, "João": 2000}
	salarios["Anderson"] = 5000

	for c, v := range salarios {
		fmt.Printf(" O salario de %s é %d\n", c, v)
	}

	for _, salario := range salarios {
		fmt.Printf(" O salario  é %d\n", salario)
	}

}

func sum(a, b int) (int, error) {
	if a+b < 0 {
		return 0, errors.New("A soma resulta em um numero negativo")
	}
	return a + b, nil
}

func sumVar(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
