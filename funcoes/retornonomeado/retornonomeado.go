package main

import "fmt"

func main() {
	p2, p1 := trocar(2, 3)
	fmt.Println(p1, p2)
}

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo, retorna a propria variavel declarada no retorno .
}
