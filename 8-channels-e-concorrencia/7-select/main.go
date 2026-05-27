package main

import (
	"fmt"
	"time"
)

func enviarMensagem(ch chan<- string, mensagem string, intervalo time.Duration) {
	for {
		time.Sleep(intervalo)
		ch <- mensagem
	}
}

func main() {
	var (
		ch1 = make(chan string)
		ch2 = make(chan string)
	)

	go enviarMensagem(ch1, "Enviando mensagem para o channel 1", time.Second*3)
	go enviarMensagem(ch2, "Enviando mensagem para o channel 2", time.Second*4)

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(time.Second * 5):
			fmt.Println("Nenhuma mensagem recebida")
			return
		}
	}
}
