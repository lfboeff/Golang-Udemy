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
	Idade uint8  `json:"idade"`
}

func main() {

	c := cachorro{
		"Rex",
		"Dálmata",
		3,
	}
	fmt.Println("struct:\t", c)

	cachorroEmJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("json:\t", bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome":  "Tóbi",
		"raca":  "Linguicinha",
		"idade": "10",
	}
	fmt.Println("map:\t", c2)

	cachorro2EmJSON, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("json:\t", bytes.NewBuffer(cachorro2EmJSON))

}
