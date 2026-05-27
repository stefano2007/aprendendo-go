package main

import (
	"fmt"
	"time"
)

func agendar(intervalo time.Duration, f func(), channelDeParada <-chan struct{}, channelDeFinalizacao chan<- struct{}) {
	ticker := time.NewTicker(intervalo)
	defer close(channelDeFinalizacao)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			f()
		case <-channelDeParada:
			fmt.Println("Sinal de parada recebido!")
			return
		}
	}
}

func main() {
	var channelDeParada = make(chan struct{})      // stop
	var channelDeFinalizacao = make(chan struct{}) // done

	go agendar(time.Second*3, func() {
		fmt.Println("Executando agendamento!")
	}, channelDeParada, channelDeFinalizacao)

	time.Sleep(time.Second * 10)

	close(channelDeParada)

	//Aguarda o channel fechar, e garante que o ticker.Stop() seja chamado antes de finalizar a aplicação
	<-channelDeFinalizacao
}
