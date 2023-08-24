/*
	sprintf funciona como printf
	retorna uma string como entrada
	nao imprime na tela
*/
package main

import "fmt"

func main() {
	var nome string = "Luiz"

	str := fmt.Sprintf("Bem vindo %s", nome)
	fmt.Println(str)
}
