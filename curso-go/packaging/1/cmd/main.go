package main

import (
	"fmt"
	"github.com/KlausBarbosa/packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.C)
	//m := math.Math{A: 1, B: 2}
	//fmt.Println(m.Add())
}
