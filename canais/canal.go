package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo", canal)
	fmt.Println("Depois que a função escrever, começa a execução")
	/*
		for {
			mensagem, open := <-canal
			if !open { // verifica se o canal está fechado ou aberta, fechado, ele sai do for
				break
			}
			fmt.Println(mensagem)
		}
	*/

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal) // fechando um canal
}
