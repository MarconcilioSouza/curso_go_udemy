package main

import "fmt"

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

type pedido struct {
	numero  int
	Produto Produto
}

// Método: função com receiver (receptor)
func (p Produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	var Produto1 Produto
	Produto1 = Produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(Produto1)

	deconto := Produto1.precoComDesconto()
	fmt.Println(deconto)

	var pedido1 pedido
	pedido1 = pedido{
		numero:  1,
		Produto: Produto1,
	}

	fmt.Println(pedido1.Produto.precoComDesconto())
	fmt.Println(pedido1.Produto.nome)

	produto2 := Produto{"Caderno", 2.45, 0.09}
	fmt.Println(produto2)
}
