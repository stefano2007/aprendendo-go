package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Map
		wg sync.WaitGroup
	)
	const numGoroutines = 100

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			m.Store(i, i)
		}()
	}

	wg.Wait()

	m.Range(func(chave, valor any) bool {
		fmt.Println(chave, valor)
		return true
	})

	valor, ok := m.Load(50)
	if ok {
		fmt.Println("Valor encontrado", valor)
	} else {
		fmt.Println("Nao encontrado")
	}

}
