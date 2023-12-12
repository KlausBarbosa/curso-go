package main

import (
	"github.com/KlausBarbosa/packaging/3/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())

}
