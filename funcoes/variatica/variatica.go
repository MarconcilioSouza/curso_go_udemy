package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	for _, valor := range numeros {
		total += valor
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Media: %.2f", media(3.2, 5.9, 9.8, 9.6))
}
