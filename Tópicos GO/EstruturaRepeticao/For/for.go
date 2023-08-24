package main

import "fmt"

func main() {
	var (
		idade int
		pares int
		nome  string
	)

	fmt.Print("Informe seu nome e sua idade: ")
	fmt.Scan(&nome, &idade)

	for i := 0; i < idade; i += 2 {
		pares++
	}

	fmt.Println(nome, "existem", pares, "numeros pares ate sua idade.")
}
