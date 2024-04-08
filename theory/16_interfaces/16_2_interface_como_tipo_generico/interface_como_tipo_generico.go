package main

import "fmt"

type interfaceGenerica interface{} // criamos uma interface como tipo genérico a partir da criação de uma interface vazia (interface{})

func escreveFormaGenerica(i interface{}) {
	fmt.Println(i)
}

type retangulo struct {
	nome    string
	altura  float64
	largura float64
}

type circulo struct {
	nome string
	raio float64
}

func main() {

	r1 := retangulo{"retângulo", 10, 15}
	escreveFormaGenerica(r1)

	c1 := circulo{"círculo", 10}
	escreveFormaGenerica(c1)

	int1 := 10
	escreveFormaGenerica(int1)

	str1 := "Felipe"
	escreveFormaGenerica(str1)

	bool1 := true
	escreveFormaGenerica(bool1)

	map1 := map[interface{}]interface{}{
		1:            "string1",
		float32(100): true,
		"string2":    "string3",
		true:         float64(5),
	}
	fmt.Println(map1)

}
