package main

import "math"

// Iniciando com letra maiúscula é PÚBLICO (visivel dentro e fora pacote)!
// Iniciando com letra minúscula é PRIVADO (visivel apenas dentro do pacote)

// Por exemplo...
// "Ponto" - gerará algo público
// "ponto" ou "_ponto" - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

// função privado por iniciar com minúscula
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return
}

// funcção publica por iniciar com maiúscula
// Distancia é responsável por calcula a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
