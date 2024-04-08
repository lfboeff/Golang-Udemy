package main

import (
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
	cachorroEmJson := `{"nome":"Rex","raca":"Dálmata","idade":3}`
	fmt.Println(cachorroEmJson)

	var c cachorro // ou c := cachorro {}
	fmt.Println(c)

	if err := json.Unmarshal([]byte(cachorroEmJson), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	fmt.Println("---")

	cachorroEmJson2 := `{"idade":"10","nome":"Tóbi","raca":"Linguicinha"}`
	fmt.Println(cachorroEmJson2)

	c2 := make(map[string]string)
	fmt.Println(c2)

	if err := json.Unmarshal([]byte(cachorroEmJson2), &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)

}
