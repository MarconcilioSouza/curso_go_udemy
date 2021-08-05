package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas deve ser inicializados

	aprovados := make(map[int]string)

	aprovados[1232434324] = "Maria"
	aprovados[1232434323] = "Pedro"
	aprovados[1232434322] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[1232434322])
	delete(aprovados, 1232434322)
}
