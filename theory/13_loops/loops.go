package main

import (
	"fmt"
)

func main() {

	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrimentando i:", i)
		// time.Sleep(time.Second)
	}

	fmt.Println()
	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j:", j)
		// time.Sleep(time.Second)
	}

	fmt.Println()
	nomes := []string{"Luís", "Berwanger", "Hanauer"}
	for index, nome := range nomes {
		fmt.Println(index, "-", nome)
	}

	fmt.Println()
	palavra := "Aloha irmãos"
	for index, letra := range palavra {
		fmt.Println(index, "-", string(letra))
	}

	fmt.Println()
	usuario := map[string]string{
		"nome":     "Felipe",
		"idade":    "33",
		"país":     "Brasil",
		"telefone": "(51) 9 9999 9999",
	}
	for key, value := range usuario {
		fmt.Println(key, "-", value)
	}

	// o comando "range" não funciona para structs

}
