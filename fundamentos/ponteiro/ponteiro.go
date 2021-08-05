package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritimética de ponteiros
	var p *int = nil

	p = &i // pegando o endereço da variável
	*p++   // desreferenciado (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
