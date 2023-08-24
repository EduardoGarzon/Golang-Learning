/*
	nao le \n nem concatena \n no final
*/
package main

import "fmt"

func main() {
	var nome string
	var idade int

	fmt.Scan(&idade, &nome)

	fmt.Println(nome)
	fmt.Print(idade)
}