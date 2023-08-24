package main

import (
	"fmt"
)

func main() {
	var numero int

	// Definicao de uma funcao anonima
	// Funcao anonima acionada em relacao a variavel total
	// nao precisa de nome
	// funcao restrita ao escopo da fucao que a contem
	total := func(num int) (pares int, impares int) {
		for i := 0; i < num; i++ {
			if i%2 == 0 {
				pares++
			} else if i%2 != 0 {
				impares++
			}
		}

		return pares, impares
	}

	fmt.Print("Informe um valor: ")
	fmt.Scan(&numero)
	pares, impares := total(numero)
	fmt.Println("Total de pares: ", pares)
	fmt.Println("Total de impares: ", impares)
}
