package main

import (
	"fmt"
)

func main() {
	alunosAprovados := []string{"Maria", "João", "Rafael", "Guilherme"} // declaração de um slice

	imprimirAprovados(alunosAprovados...) // os tres ... faz com que seja todos os valores sejam separados um a um
}

func imprimirAprovados(aprovados ...string) { // forma de declara uma função que recebe N parametros string
	fmt.Println("Lista de aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}
