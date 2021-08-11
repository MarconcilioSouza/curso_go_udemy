package main

import "fmt"

func rotina(ch chan int) {
	fmt.Println("Executou!")
	ch <- 1
	ch <- 2
	ch <- 3 // até aqui tem espaço no buffer
	ch <- 4
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) // 3 é o total de buffer

	go rotina(ch)

	fmt.Println(<-ch)
}
