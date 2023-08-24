/*
	funciona como println
	retorna uma string de entrada
	nao imprime na tela
*/
package main

import "fmt"

func main() {
	nome := "Luiz"

	str := fmt.Sprintln("Bem vindo", nome)
	fmt.Print(str)
	fmt.Print("Nova linha")
}
