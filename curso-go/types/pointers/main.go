package main

func main() {
	a := 19
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	println(*b)

}
