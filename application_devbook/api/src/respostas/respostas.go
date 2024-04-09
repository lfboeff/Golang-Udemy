package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em json para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {

	w.Header().Set("Content-Type", "application/json") // faz com que o client que realizou a request (Postman) receba os dados no formato JSON diretamente

	w.WriteHeader(statusCode)

	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}

}

// Erro retorna um erro em formato json
func Erro(w http.ResponseWriter, statusCode int, err error) {

	JSON(w, statusCode, struct {
		Erro string `json:"error"`
	}{
		Erro: err.Error(),
	})

}
