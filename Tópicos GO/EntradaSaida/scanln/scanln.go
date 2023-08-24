/*
	Funciona igual scan
	interrompe a leitura quando encontra \n
*/
package main

import "fmt"

func main() {
	var nome string
	var idade int

	fmt.Scanln(&nome, &idade)
	str := fmt.Sprintln("Bem vindo", nome, "voce tem", idade, "anos de idade.")
	fmt.Print(str)
}
