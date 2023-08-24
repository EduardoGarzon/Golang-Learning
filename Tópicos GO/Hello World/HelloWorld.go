/* Exemplo de programa em Go
Compilacao: go build nome.go
Compilacao e Execucao: go run nome.go
*/

// comentarios em Go seguem o mesmo padrao C/C++

/*
nome do pacote criado, equivale ao nome do programa,
indica que é executavel ao compilador,
define o ponto de entrada de um programa Go,
apenas o package "main" pode ter a função "main".
*/
package main

import "fmt" // importacao individual

// importacao de bibliotecas, equivale ao #include em C/C++
// ftm possui funcoes basicas de formatacao para entrada/saida

// ponto de partida
func main() { // identacao das chaves obrigatoria
	// comandos

	// println
	// adiciona automaticamente um caractere de nova linha (\n) ao final da saída.
	fmt.Println("Hello World!")
	fmt.Println("Hello", "World!")

	// print
	// strings para o console sem adicionar uma nova linha.
	fmt.Print("Hello World!")
	fmt.Print("Hello", "World!\n")

	// printf
	// string formatada para o console
	fmt.Printf("Hello World\n")
	fmt.Printf("Valor: %d, String: %s\n", 10, "Golang")
}
