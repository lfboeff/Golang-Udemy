package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint8  `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`
	fmt.Println("json:\t", cachorroEmJSON)

	var c cachorro // ou c := cachorro{}
	// fmt.Println(c)

	if err := json.Unmarshal([]byte(cachorroEmJSON), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println("struct:\t", c)

	cachorro2EmJSON := `{"nome":"Tóbi","raca":"Linguicinha","idade":"10"}`

	var c2 = map[string]string{} // ou c2 : make(map[string]string)
	// fmt.Println(c2)

	if err := json.Unmarshal([]byte(cachorro2EmJSON), &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("map:\t", c2)

}

// se não quiseremos utilizar algumas das tags json, podemos declará-la na struct como `json:"-"`
// existe uma função chamada "decode" que pode ser usada além de Marshal e Unmarshal > veremos melhor na aplicação
