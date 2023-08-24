package main

import (
	"fmt"
	"os"
	//"strings"
)

var (
	op       int
	nome     string
	numero   int64
	contatos = make([]string, 0, 5)
	numeros  = make([]int64, 0, 5)
)

func Menu() {
	fmt.Println("[1] Cadastrar Contatos")
	fmt.Println("[2] Exibir Contatos")
	fmt.Println("[3] Deletar Contato")
	fmt.Println("[4] Encerrar")
	fmt.Print("Selecione a opcao: ")
	fmt.Scan(&op)
	fmt.Printf("\n")

	option(op)
}

func option(op int) {
	switch op {
	case 1:
		fmt.Print("Informe o nome e numero: ")
		fmt.Scan(&nome, &numero)
		contatos = append(contatos, nome)
		numeros = append(numeros, numero)
		fmt.Printf("\n")
		Menu()
	case 2:
		for i := 0; i < len(contatos); i++ {
			fmt.Printf("NOME: %s NUMERO: %d\n", contatos[i], numeros[i])
		}
		fmt.Printf("\n")
		Menu()
	case 3:
		fmt.Print("Informe o nome do contato: ")
		fmt.Scan(&nome)
		for i := 0; i < len(contatos); i++ {
			if nome == contatos[i] {
				contatos = append(contatos[:i], contatos[i+1:]...)
				numeros = append(numeros[:i], numeros[i+1:]...)
			}
		}
		fmt.Printf("\n")
		Menu()
	case 4:
		fmt.Printf("ENCERRANDO ATIVIDADE!\n")
		os.Exit(0)
	}
}

func main() {
	Menu()
}
