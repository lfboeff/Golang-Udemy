// TESTES DE UNIDADE

// go test -v -cover >> executa os testes, mostra se cada função testada passou ou não no teste, e apresenta a porcentagem de cobertura dos testes executados
// go test -v -coverprofile resultado_teste.txt >> faz a mesma coisa que o comando anterior + salva os resultados do teste no arquivo "resultado_teste.txt"
// go tool cover -func=resultado_teste.txt >> analisa os resultados salvos no arquivo "resultado_teste.txt" e informa a porcentagem de linhas testadas para cada função original
// go tool cover -html=resultado_teste.txt >> abre uma página html temporária que aponta exatamente que linha(s) não foi(ram) testada(s) (rulez!)

package enderecos_test

import (
	. "go_tests_module/19_1_1_enderecos" // o "." significa para o arquivo que esse na verdade é o pacote principal do arquivo
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	// t.Parallel() // se permitimos um teste a rodar em paralelo, temos que adicionar isso na(s) outra(s) funcão(ões) também, ou então isso não terá nenhum efeito

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada Qualquer", "Estrada"},
		{"Praça das Rosas", "Tipo de endereço inválido"},
		{"RUA DOS BOBOS", "Rua"},
		{"rua dos bobos", "Rua"},
		{"AVENIDA", "Avenida"},
		{"avenida", "Avenida"},
		{"av", "Tipo de endereço inválido"},
		{"ESTRAD COLONIA", "Tipo de endereço inválido"},
		{"", "Tipo de endereço inválido"},
		{" ", "Tipo de endereço inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado: %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {

	// t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}

// AULA ANTERIOR:
// func TestTipoDeEndereco(t *testing.T) {
// 	// "Test" + "function name to be tested" (t *testing.T) >> best practice

// 	// enderecoParaTeste := "Avenida Paulista"
// 	enderecoParaTeste := "Rua ABC"
// 	tipoDeEnderecoEsperado := "Avenida"
// 	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
// 		t.Errorf("O tipo de endereço recebido é diferente do tipo esperado! Esperava %s e recebeu %s.\n",
// 			tipoDeEnderecoEsperado,
// 			tipoDeEnderecoRecebido)
// 	}
// }
