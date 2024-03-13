package main

import (
	"fmt"
	"math"
)

type forma interface { // interfaces só possuem assinaturas de métodos, nada mais
	getArea() float64
	getNome() string
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma %v é %0.2f\n", f.getNome(), f.getArea())
}

type retangulo struct {
	nome    string
	altura  float64
	largura float64
}

func (r retangulo) getArea() float64 {
	return r.altura * r.largura
}

func (r retangulo) getNome() string {
	return r.nome
}

type circulo struct {
	nome string
	raio float64
}

func (c circulo) getArea() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (c circulo) getNome() string {
	return c.nome
}

func main() {
	r1 := retangulo{"retângulo", 10, 15}
	escreverArea(r1)

	c1 := circulo{"círculo", 10}
	escreverArea(c1)

}
