package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Arrays: lista de valores com tamanho fixo

	var arr1 [5]int
	fmt.Println(arr1)

	arr1[0], arr1[1], arr1[2], arr1[3], arr1[4] = 5, 6, 7, 8, 9
	fmt.Println(arr1)

	fmt.Println()
	var arr2 [5]string
	fmt.Println(arr2)

	arr2[0], arr2[1], arr2[2], arr2[3], arr2[4] = "Hey", "Ho", "Let's", "Go", "!"
	fmt.Println(arr2)

	fmt.Println()
	arr3 := [...]int{1, 2, 3, 4, 5} // os "..." fazem com que o array seja criado conforme o número de itens entre "{}"
	fmt.Println(arr3)

	// Slices: lista de valores com tamanho variável

	fmt.Println()
	var slice1 = []int{0, 0, 1, 3, 5}
	fmt.Println(slice1)

	slice1[0] = 1
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(arr3))
	fmt.Println(reflect.TypeOf(slice1))
	// como o tipo do array é diferente do tipo do slice, não podemos fazer operações diretamente entre os mesmos

	fmt.Println()
	fmt.Println(slice1)
	slice1 = append(slice1, 9, 8, 7)
	fmt.Println(slice1)

	fmt.Println()
	slice2 := arr1[1:3]
	fmt.Println(slice2)

	arr1[1] = 95 // quando alteramos o array interno de slice2, o valor correspondente em slice2 acaba sendo alterado também
	fmt.Println(slice2)

	// Arrays internos:

	fmt.Println("\n----- Arrays Internos:")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 7)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println()
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 4)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
