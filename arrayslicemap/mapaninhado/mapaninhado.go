package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15465.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)

		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
