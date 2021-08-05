package main

import "fmt"

func main() {
	funcsESalario := map[string]float64{
		"José João": 1122.21,
		"Maria":     2323.21,
		"Pedro":     4332.21,
		"paulo":     7899.21,
	}

	funcsESalario["José João"] = 54245.85

	delete(funcsESalario, "Inexistente") // teste para ver se dar erro

	for nome, salario := range funcsESalario {

		fmt.Println(nome, salario)
	}
}
