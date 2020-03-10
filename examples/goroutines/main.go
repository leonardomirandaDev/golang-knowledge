package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hardRoutine(name string) {
	// defer chamará a wg.Done quando a função onde ele está inserido acabar
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		time.Sleep(1 + time.Second)
		fmt.Printf("Executando tarefa %s \n", name)
	}
	fmt.Printf("\nTarefa %s finalizada\n\n", name)
}

func main() {
	for i := 1; i <= 10; i++ {
		// adiciona um contador de funcoes que estao sendo executadas
		wg.Add(1)
		go hardRoutine(strconv.Itoa(i))
	}

	// espera até nao existir funcoes no contador wg
	wg.Wait()
	fmt.Printf("\nTarefas finalizadas\n")
}
