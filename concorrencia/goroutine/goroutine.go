package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Por quê você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// o go faz com que seja chamada a função sem que ela espere o processamento
	go fale("Maria", "Por quê você não fala comigo?", 3)
	go fale("João", "Só posso falar depois de você!", 1)
}
