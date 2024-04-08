package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// json.Marshal
	c := cachorro{
		Nome:  "Rex",
		Raca:  "Dálmata",
		Idade: 3,
	}
	fmt.Println(c)

	cachorroEmJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cachorroEmJson) // saída igual a um slice de bytes: []byte ([]uint8)

	// para resolver isso e assim vermos um JSON no formato normal, podemos usar a libraby "bytes":
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	fmt.Println("---")

	c2 := map[string]string{
		"nome":  "Tóbi",
		"raca":  "Linguicinha",
		"idade": "10",
	}
	fmt.Println(c2)

	cachorroEmJson2, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cachorroEmJson2)
	fmt.Println(bytes.NewBuffer(cachorroEmJson2))

}
