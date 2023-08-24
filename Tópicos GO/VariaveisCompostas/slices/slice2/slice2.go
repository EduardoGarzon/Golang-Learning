package main

import (
	"fmt"
)

func main() {
	var (
		pessoas int
		peso    int
		nome    string
	)

	fmt.Print("Informe o total de individuos: ")
	fmt.Scanf("%d%*c", &pessoas)

	// funcao embutida make (tipo, comprimento, capacidade)
	individuos := make([]string, 0, 3)
	pesos := make([]float64, 0, 3)

	for i := 0; i < pessoas; i++ {
		fmt.Printf("Informe o nome e peso do individuo [%d]: ", i+1)
		fmt.Scan(&nome, &peso)

		if len(individuos) < cap(individuos) {
			individuos = append(individuos, nome)
			pesos = append(pesos, float64(peso))
		} else {
			fmt.Println("CAPACIDADE MAXIMA ATINGIDA!")

		}
	}

	fmt.Println("Individuos Cadastrados")
	fmt.Println("NOMES: ", individuos)
	fmt.Println("PESOS: ", pesos)
}
