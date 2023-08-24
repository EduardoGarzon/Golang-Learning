package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	LIN = 3
	COL = 3
)

func main() {
	var (
		A            [LIN][COL]float64
		B            [LIN][COL]float64
		somaA, somaB float64
	)

	// Inicializa o gerador de números aleatórios com a hora atual como semente
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Gera um valor float64 aleatório no intervalo [0.0, 1.0) e o escala para o intervalo [0.0, 100.0)
			A[i][j] = rand.Float64() * 100
			B[i][j] = rand.Float64() * 100
			somaA += A[i][j]
			somaB += B[i][j]
		}
	}

	fmt.Println("MATRIZ A:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%0.2f ", A[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Println("MATRIZ B:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%0.2f ", B[i][j])
		}
		fmt.Printf("\n")
	}

	if somaA > somaB {
		fmt.Printf("A soma dos valores da matriz A é maior: %0.2f\n", somaA)
	} else {
		fmt.Printf("A soma dos valores da matriz B é maior: %0.2f\n", somaB)
	}

}
