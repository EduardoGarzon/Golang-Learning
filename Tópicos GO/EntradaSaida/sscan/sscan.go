/*
	sscan le a entrada a partir de uma string de argumento
	funciona de forma semelhante ao scan
*/
package main

import "fmt"

func main() {
	var idade int
	var valor string

	fmt.Scan(&valor)
	fmt.Sscan(valor, &idade)
	fmt.Println("String informada: ", idade)
}
