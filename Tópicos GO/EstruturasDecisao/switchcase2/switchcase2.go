package main

import "fmt"

func main() {
	var nome string

	fmt.Scanf("%s%*c", &nome)

	switch nome {
	case "luiz":
		nome = fmt.Sprint(nome + " Garzon")
	case "eduarda":
		nome = fmt.Sprint(nome + " Kalschne")
	default:
		fmt.Println("Erro, parametros invalidos.")
	}
	fmt.Println(nome)
}
