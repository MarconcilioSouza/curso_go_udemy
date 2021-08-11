package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	fmt.Println("Antes de lê")

	c <- 1 // operação bloqueio
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer
	go rotina(c)
	fmt.Println(<-c)
	fmt.Println("Foi lido")

	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")
}
