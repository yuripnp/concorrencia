package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// primeiro método para sincronizar gorotines
	var waitgroup sync.WaitGroup // um grupo de espera

	waitgroup.Add(2) // quantidade de gorotines que está dentro do grupo de espera

	go func() {
		escrever("Olá mundo")
		waitgroup.Done() // -1
	}()

	go func() {
		escrever("go rotinas go")
		waitgroup.Done() // -1
	}()

	waitgroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
