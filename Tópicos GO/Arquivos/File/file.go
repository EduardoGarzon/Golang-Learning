package main

import (
	"fmt"
	"os"
)

func main() {
	// Tenta abrir o arquivo.go ja existente
	file, ferror := os.ReadFile("../HelloWorld.go")

	// Testa se ocorreu erro
	if ferror != nil { // nil é semelhante ao NULL
		fmt.Println("Erro: ", ferror) // imprime o erro ocorrido
		os.Exit(1)                    // encerra execucao
	} else {
		file2, ferror := os.Create("arquivo.txt") // Cria e abre o arquivo.txt
		if ferror != nil {
			fmt.Println("Erro: ", ferror)
			os.Exit(1)
		}
		content := string(file)    // processa o arquivo.go
		file2.WriteString(content) // escreve no arquivo.txt
		fmt.Println(content)

		file2.Close()
	}
}
