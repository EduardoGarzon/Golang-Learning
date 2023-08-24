package main

// importacao em conjunto
import (
	"fmt"
)

/*	declaracao multipla de variaveis globais em bloco,
	declaracao de variaveis em massa ou em grupo.
*/

var (
	booleano bool // um valor booleano que pode ser verdadeiro (true) ou falso (false).

	caracteres string = "Luiz" // uma sequência de caracteres Unicode.

	inteiro   int   = 0   // max 2^63 - 1 (9223372036854775807) e min -2^63 (-9223372036854775808)
	inteiro8  int8  = 20  // max 127 e min -128
	inteiro16 int16 = 30  // max 32767 e min -32768
	inteiro32 int32 = -40 // max 32767 e min -32768.
	inteiro64 int64 = -50 // max 9223372036854775807 e min -9223372036854775808.

	uinteiro   uint   = 0   // max 2^64 - 1 (18446744073709551615) e min 0.
	uinteiro8  uint8  = 70  // max 255 e min 0.
	uinteiro16 uint16 = 80  // max 65535 e min 0.
	uinteiro32 uint32 = 90  // max 4294967295 e min 0.
	uinteiro64 uint64 = 100 // max 18446744073709551615 e min 0

	uintponteiro uintptr = 18446744073709551615 // inteiro sem sinal grande o suficiente para armazenar o valor de um ponteiro ou endereço de memória.

	/*
		Em resumo, um alias é uma nova declaração de um tipo existente, permitindo que ele seja referenciado por um nome diferente
	*/
	c byte = 255 // um alias para uint8, dados de caracteres em forma numerica, único byte de dados.

	r1, r2 rune = '1', '2' // um alias para int32 que representa um ponto de código Unicode, caracteres UTF-8.

	f1 float32 = -1.1234    // max 3.4x10^38 e min -3.4x10^38.
	f2 float64 = 1.79763134 // max 1.8x10^308 e min -1.8x10^308

	com1 complex64  = 9 + 10i  // qualquer valor dentro do intervalo dos valores float32.
	com2 complex128 = 10 + 10i // qualquer valor dentro do intervalo dos valores float64.
)

func main() {
	// declaracao de variaveis do tipo short hand (somente dentro de funcoes)
	// o tipo é inferido automaticamente pelo valor atribuido
	i1 := 'A'     // rune ou char
	i2 := "texto" // string
	i3 := 2       // int
	i4 := 2.5     // float

	// variaveis globais
	fmt.Println("Variaveis Globais:")
	fmt.Printf("Tipo: %T | Valor: %v\n", booleano, booleano)
	fmt.Printf("Tipo: %T | Valor: %v\n", caracteres, caracteres)

	fmt.Printf("Tipo: %T | Valor: %v\n", inteiro, inteiro)
	fmt.Printf("Tipo: %T | Valor: %v\n", inteiro8, inteiro8)
	fmt.Printf("Tipo: %T | Valor: %v\n", inteiro16, inteiro16)
	fmt.Printf("Tipo: %T | Valor: %v\n", inteiro32, inteiro32)
	fmt.Printf("Tipo: %T | Valor: %v\n", inteiro64, inteiro64)

	fmt.Printf("Tipo: %T | Valor: %v\n", uinteiro64, uinteiro64)
	fmt.Printf("Tipo: %T | Valor: %v\n", uinteiro64, uinteiro64)
	fmt.Printf("Tipo: %T | Valor: %v\n", uinteiro64, uinteiro64)

	fmt.Printf("Tipo: %T | Valor: %v\n", uintponteiro, uintponteiro)

	fmt.Printf("Tipo: %T | Valor: %v\n", c, c)

	fmt.Printf("Tipo: %T | Valor: %v\n", r1, r1)
	fmt.Printf("Tipo: %T | Valor: %v\n", r2, r2)

	fmt.Printf("Tipo: %T | Valor: %v\n", f1, f1)
	fmt.Printf("Tipo: %T | Valor: %v\n", f2, f2)

	fmt.Printf("Tipo: %T | Valor: %v\n", com1, com1)
	fmt.Printf("Tipo: %T | Valor: %v\n", com2, com2)

	// variaveis locais
	fmt.Println()
	fmt.Println("Variaveis Locais:")
	fmt.Printf("Tipo: %T | Valor: %v\n", i1, i1)
	fmt.Printf("Tipo: %T | Valor: %v\n", i2, i2)
	fmt.Printf("Tipo: %T | Valor: %v\n", i3, i3)
	fmt.Printf("Tipo: %T | Valor: %v\n", i4, i4)

}
