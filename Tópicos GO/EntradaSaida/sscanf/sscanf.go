/*
	sscanf le a entrada a partir de uma string de argumento formatada
	semelhante ao scanf
*/

package main

import "fmt"

func main() {
	var dados string = "Nome: Luiz Idade: 20"
	var nome string
	var idade int

	fmt.Sscanf(dados, "Nome: %s Idade: %d", &nome, &idade)
	fmt.Println("Nome:", nome, "Idade:", idade)
}
