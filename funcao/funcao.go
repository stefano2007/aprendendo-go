package main

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Soma(a, b int) int {
	return a + b
}

// modifica o parametro passado por referencia
func raizQuadrada(num *float64) {
	aux := *num
	*num = aux * aux
}

func main() {
	messageOutput := Hello("Golang")
	fmt.Println(messageOutput)

	a, b := 10, 20

	result := Soma(a, b)
	fmt.Printf("Soma de (%v + %v) = %v \n", a, b, result)

	valor := 5.0
	valorOriginal := valor
	raizQuadrada(&valor)
	fmt.Printf("raiz quadrada de (%v) = %v \n", valorOriginal, valor)
}
