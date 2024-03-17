package enderecos

import "testing"

// TESTES DE UNIDADE

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

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
