package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero_int16 int16 = 32_767
	fmt.Println(numero_int16)

	var numero_pos uint16 = 500
	fmt.Println(numero_pos)

	// aliases:

	// int32 = rune
	var numero_int32 rune = 2_147_483_647
	fmt.Println(numero_int32)

	// uint8 = byte
	var numero_uint8 byte = 255
	fmt.Println(numero_uint8)

	var numero_real_1 float32 = 123.45
	var numero_real_2 float64 = 123.45
	fmt.Println(numero_real_1, numero_real_2)

	numero_real_3 := 123.45
	fmt.Printf("%v - %T\n", numero_real_3, numero_real_3)

	var str_1 string = "Texto"
	fmt.Println(str_1)

	str_2 := "Texto 2"
	fmt.Println(str_2)

	char_1 := 'A'
	fmt.Println(char_1) // declarar um caractere com aspas simples, faz com que o valor salvo seja o inteiro que o representa na tabela ASCII

	char_2 := "A"
	fmt.Println(char_2)

	var texto string
	fmt.Println("string \"valor zero\":", texto)

	var inteiro int
	fmt.Println("int \"valor zero\":", inteiro)

	var bool_1 bool
	fmt.Println("bool \"valor zero\":", bool_1)

	var error_1 error
	fmt.Println("error \"valor zero\":", error_1)

	var error_2 error = errors.New("Erro interno")
	fmt.Println(error_2)
}
