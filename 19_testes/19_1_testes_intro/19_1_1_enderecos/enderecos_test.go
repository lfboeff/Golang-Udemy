package enderecos

import "testing"

// TESTES DE UNIDADE

func TestTipoDeEndereco(t *testing.T) {
	// "Test" + "function name to be tested" (t *testing.T) >> best practice

	// enderecoParaTeste := "Avenida Paulista"
	enderecoParaTeste := "Rua ABC"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo de endereço recebido é diferente do tipo esperado! Esperava %s e recebeu %s.\n",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido)
	}

}
