package main

import "fmt"

type mySlice []int

func main() {
	var lista = mySlice{1, 2, 3, 4, 5, 6}

	lista = lista.Filter(func(i int) bool {
		return i%2 == 0
	})

	fmt.Println(lista)

	lista = lista.Map(func(i int) int {
		return i * 10
	})

	fmt.Println(lista)

	mult := lista.Reduce(func(i1, i2 int) int {
		return i1 * i2
	}, 1)

	fmt.Println(mult)

	fmt.Println("----------")

	var lista2 = mySlice{1, 2, 3, 4, 5, 6}

	mult2 := lista2.Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i * 10
	}).Reduce(func(i1, i2 int) int {
		return i1 * i2
	}, 1)

	fmt.Println(mult2)
}

func (m mySlice) Filter(cond func(int) bool) mySlice {
	var resultado mySlice

	for _, numero := range m {
		if cond(numero) {
			resultado = append(resultado, numero)
		}
	}

	return resultado
}

func (m mySlice) Map(trans func(int) int) mySlice {
	var resultado mySlice

	for _, numero := range m {
		resultado = append(resultado, trans(numero))
	}

	return resultado
}

func (m mySlice) Reduce(acc func(int, int) int, inicial int) int {
	var resultado = inicial

	for _, numero := range m {
		resultado = acc(resultado, numero)
	}

	return resultado
}
