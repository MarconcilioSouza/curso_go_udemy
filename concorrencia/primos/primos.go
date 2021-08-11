package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 0; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			c <- primo
			inicio = inicio + 1
			time.Sleep(time.Millisecond * 100)
			break
		}
	}
	close(c) // finaliza o canal
}

func main() {
	c := make(chan int, 30)

	go primos(cap(c), c)

	for primo := range c {
		fmt.Printf("%d\n", primo)
	}

	fmt.Println("Fim")
}
