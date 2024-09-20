package main

import "fmt"

func main() {
	canal := make(chan string, 1) // buffer -> coloca a capacidade desse canal

	canal <- "Olá mundo"
	mensagem := canal

	fmt.Println(mensagem)
}
