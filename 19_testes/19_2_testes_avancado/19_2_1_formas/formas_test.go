// TDD: TEST DRIVEN DEVELOPMENT >> os testes são escritos antes das funções originais, as quais são então criadas e testadas aos poucos

package formas

import (
	"math"
	"testing"
)

func TestGetArea(t *testing.T) {

	// como temos dois métodos a serem testados, podemos usar "t.Run()" para testá-los separadamente
	t.Run("Teste do Retângulo", func(t *testing.T) {
		ret := Retangulo{Altura: 10, Largura: 12}
		areaEsperada := float64(120)
		areaRecebida := ret.GetArea()

		if areaRecebida != areaEsperada {
			// t.Errorf("A área recebida %v difere da área esperada: %v", areaRecebida, areaEsperada)
			t.Fatalf("A área recebida do retâgulo (%v) difere da área esperada (%v)", areaRecebida, areaEsperada)
			// o "t.Fatalf" fará com que os testes subsequentes não sejam mais executados, ao contrário do que aconteceria com "t.Errorf"
		}

	})

	t.Run("Teste do Círculo", func(t *testing.T) {
		circ := Circulo{Raio: 10}
		areaEsperada := float64(math.Pi * 10 * 10)
		areaRecebida := circ.GetArea()

		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida do círculo (%v) difere da área esperada (%v)", areaRecebida, areaEsperada)
		}
	})
}

// go test -v -cover
// go test -v -coverprofile resultado_teste.txt
// go tool cover -html=resultado_teste.txt
