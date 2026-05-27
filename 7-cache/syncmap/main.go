package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	dados sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(chave string, valor any) {
	c.dados.Store(chave, valor)
	fmt.Printf("Valor %v armazenado no cache com a chave %s\n", valor, chave)
}

func (c *Cache) Get(chave string) (any, bool) {
	return c.dados.Load(chave)
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(fmt.Sprintf("chave_%d", i), i)
		}()
	}

	time.Sleep(100 * time.Millisecond)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			valor, ok := cache.Get(fmt.Sprintf("chave_%d", i))
			if ok {
				fmt.Printf("Valor encontrado na chave_%d: %v\n", i, valor)
			} else {
				fmt.Printf("Valor nao encontrado na chave_%d\n", i)
			}
		}()
	}

	wg.Wait()
}
