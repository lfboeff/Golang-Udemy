package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEndereco verifica se um endereço possui um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {

	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0] // pegamos apenas a primeira palavra de "endereco"

	enderecoTemUmTipoValido := false
	for _, tipoValido := range tiposValidos {
		if tipoValido == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return cases.Title(language.Und, cases.NoLower).String(primeiraPalavraDoEndereco)
	}
	return "Tipo de endereço inválido"

}
