package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "Joao": 2000, "Anderson": 3000}
	m2 := map[string]float64{"Wesley": 1000.10, "Joao": 2000.30, "Anderson": 3000.55}
	println(Soma(m))
	println(Soma(m2))
	println(Compara(10, 10.0))
}
