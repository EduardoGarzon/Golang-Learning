package main

import "fmt"

func main() {
	nome := "Luiz"
	
	fmt.Print("Bem vindo " + nome + "\n")
	
	fmt.Print("Bem vindo ", nome, "\n")

	fmt.Print("Bem vindo\n")
	fmt.Print(nome)
}