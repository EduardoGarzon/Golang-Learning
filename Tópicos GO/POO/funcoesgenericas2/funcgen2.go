package main

import "fmt"

// a interface indica quais os tipos podem ser utilizados na funcao
type TipoGenerico interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64 | complex64 | complex128
}

// a funcao implementa o tipo generico (T) da interface que contem os tipos listados
func somaGenerica[T TipoGenerico](a T, b T) T {
	return a + b
}

func main() {
	// pode-se declarar variaveis com tipos diferentes
	// nao se pode misturar variaveis de tipos diferentes pois go nao faz cast implicito
	a := 2.5
	b := 3.5
	fmt.Println("Soma para tipos Genericos:", somaGenerica(a, b))
}
