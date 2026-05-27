package main

import (
	"fmt"
	"time"
)

func consumidor(id int, ch <-chan int) {
	for valor := range ch {
		fmt.Printf("Consumidor %d recebeu a mensagem: %d\n", id, valor)
	}
}

func broadcast(channels []chan int, valor int) {
	for _, ch := range channels {
		ch <- valor
	}
}

func main() {
	var channels = []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}

	for i, ch := range channels {
		go consumidor(i+1, ch)
	}

	broadcast(channels, 100)
	time.Sleep(time.Second)
}
