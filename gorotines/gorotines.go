package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO
	// PARALELISMO = acontece quanto tenho duas ou mais tarefas acontecendo ao mesmo tempo, isso so é possivel se tiver um processador com mais de um nucleo,
	// CONCORRENCIA = não necessariamente estão sendo processadas simultaneamente, as tarefas rezevam tempo do nucleo do processador

	go escrever("Olá mundo") // uma gorotine, que é uma função que prossegue o programa sem esperar a função terminar
	escrever("go rotinas go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
