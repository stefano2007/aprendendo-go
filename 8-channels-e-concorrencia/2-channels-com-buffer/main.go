package main

import (
	"fmt"
	"time"
)

func enviar(ch chan<- int, index int) {
	fmt.Printf("Enviando valor para o channel %d\n", index)
	ch <- 100
	fmt.Printf("Valor enviado com sucesso %d\n", index)
}

func receber(ch <-chan int, index int) {
	time.Sleep(time.Second)
	valor := <-ch
	fmt.Printf("Valor recebido %d: %d\n", index, valor)
}

func main() {
	// Comunicação entre goroutines.
	ch := make(chan int, 3) // Channel com buffer

	for i := 0; i < 5; i++ {
		go enviar(ch, i)
	}

	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		go receber(ch, i)
	}

	time.Sleep(time.Second * 5)
}
