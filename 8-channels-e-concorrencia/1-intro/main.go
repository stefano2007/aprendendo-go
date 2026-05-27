package main

import (
	"fmt"
	"time"
)

func enviar(ch chan<- int) {
	fmt.Println("Enviando valor para o channel")
	ch <- 100
	fmt.Println("Valor enviado com sucesso")
}

func receber(ch <-chan int) {
	time.Sleep(time.Second)
	valor := <-ch
	fmt.Println("Valor recebido:", valor)
}

func main() {
	// Comunicação entre goroutines.
	ch := make(chan int) // Channel sem buffer

	for i := 0; i < 3; i++ {
		go enviar(ch)
	}

	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		go receber(ch)
	}

	time.Sleep(time.Second * 5)
}
