package formas

import (
	"fmt"
	"math"
)

type Forma interface { // interfaces só possuem assinaturas de métodos, nada mais
	GetArea() float64
	GetNome() string
}

func EscreverArea(f Forma) {
	fmt.Printf("A área da forma %v é %0.2f\n", f.GetNome(), f.GetArea())
}

type Retangulo struct {
	Nome    string
	Altura  float64
	Largura float64
}

func (r Retangulo) GetArea() float64 {
	return r.Altura * r.Largura
}

func (r Retangulo) GetNome() string {
	return r.Nome
}

type Circulo struct {
	Nome string
	Raio float64
}

func (c Circulo) GetArea() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func (c Circulo) GetNome() string {
	return c.Nome
}
