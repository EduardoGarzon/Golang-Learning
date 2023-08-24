/*
	executa a leitura ate encontrar um \n
	semelhante ao scan
*/

package main

import "fmt"

func main() {
	dado := "Luiz Eduardo\n"
	var nome string
	var sobrenome string

	fmt.Sscanln(dado, &nome, &sobrenome)
	fmt.Println(nome, sobrenome)
}
