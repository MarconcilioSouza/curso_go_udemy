package main

import "fmt"

func main() {
	valor := somar(1, 3)
	imprimir(valor)
}

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println((valor))
}
