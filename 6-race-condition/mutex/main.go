package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  = make(map[int]int)
		wg sync.WaitGroup // WaitGroup é uma estrutura de sincronização que espera um conjunto de goroutines terminarem
		mu sync.Mutex     //
	)

	const numGoroutines = 100

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done() // defer é usado para garantir que wg.Done() seja chamado mesmo que ocorra um panic dentro da goroutine
			mu.Lock()       // Lock é usado para bloquear o acesso ao recurso compartilhado (mapa) para garantir que apenas uma goroutine possa acessá-lo por vez
			m[i] = i
			mu.Unlock() // Unlock é usado para liberar o bloqueio, permitindo que outras goroutines acessem o recurso compartilhado
		}()
	}

	wg.Wait() // Aguarda todas as goroutines terminarem

	for chave, valor := range m {
		fmt.Println(chave, valor)
	}
}
