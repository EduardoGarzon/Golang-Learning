package main

import "fmt"

func main() {
	var quantidade int
	var codigo int
	var valor float32

	fmt.Println("Codigo  Valor")
	fmt.Println("  0     2.50")
	fmt.Println("  1     3.00")
	fmt.Println("  2     6.00")
	fmt.Println("  3     1.50")

	fmt.Print("Informe o codigo e a quantidade: ")
	fmt.Scan(&codigo, &quantidade)

	switch codigo {
	case 0:
		valor = float32(quantidade) * float32(2.50)
	case 1:
		valor = float32(quantidade) * float32(3.00)
	case 2:
		valor = float32(quantidade) * float32(6.00)
	case 3:
		valor = float32(quantidade) * float32(1.50)
	default:
		fmt.Println("Erro, parametros invalidos.")
	}

	fmt.Printf("Valor a pagar: R$%0.2f\n", valor)
}
