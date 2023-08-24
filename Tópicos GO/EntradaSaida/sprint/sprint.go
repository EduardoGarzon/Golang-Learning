/*
	sprint funciona como print
	retorna uma string com a entrada
	nao imprime no console
	formata strings
*/
package main

import "fmt"

func main() {
	nome := "Luiz"
	idade := 20
	var str string

	str = fmt.Sprint("Seu nome é ", nome, " e voce tem ", idade, " anos de idade.")
	fmt.Println(str)
}
