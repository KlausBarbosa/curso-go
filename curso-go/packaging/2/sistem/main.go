package main

import (
	"github.com/KlausBarbosa/packaging/2/math"
	"github.com/google/uuid"
)

func main() {
	println(uuid.New().String())
	m := math.NewMath(1, 2)
	println(m.Add())
}
