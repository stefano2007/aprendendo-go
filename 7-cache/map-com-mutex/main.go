package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	dados map[string]any
	mu    sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		dados: make(map[string]any),
	}
}

func (c *Cache) Set(chave string, valor any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dados[chave] = valor
	fmt.Printf("Valor %v armazenado no cache com a chave %s\n", valor, chave)
}

func (c *Cache) Get(chave string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	valor, ok := c.dados[chave]
	return valor, ok
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(fmt.Sprintf("chave_%d", i), i)
		}()
	}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
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
