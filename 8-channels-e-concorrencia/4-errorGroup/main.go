package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d iniciado\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finalizado\n", id)
}

func workerWithErrorGroup(id int) error {
	fmt.Printf("Worker %d iniciado\n", id)

	time.Sleep(time.Second)

	if id%2 == 0 {
		return errors.New("erro inesperado id: " + fmt.Sprint(id))
	}

	fmt.Printf("Worker %d finalizado\n", id)

	return nil
}

func main() {
	fmt.Println("-------------WaitGroup-------------")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("Fim")

	fmt.Println("-------------ErrorGroup-------------")

	var eg errgroup.Group

	eg.SetLimit(3)

	for i := 1; i <= 5; i++ {
		eg.Go(func() error {
			return workerWithErrorGroup(i)
		})
	}

	err := eg.Wait()
	if err != nil {
		// imprime apenas o primeiro erro encontrado
		fmt.Println("Erro encontrado: ", err.Error())
	} else {
		fmt.Println("Fim")
	}
}
