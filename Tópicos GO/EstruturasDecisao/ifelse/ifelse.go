package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Informe sua idade: ")
	fmt.Scan(&idade)

	if idade >= 18 {
		fmt.Println("Maior de Idade.")
	} else {
		fmt.Println("Menor de Idade.")
	}
}
