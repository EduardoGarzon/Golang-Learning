package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Tenta abrir o aquivo
	file, ferror := os.Open("../HelloWorld.go") // arquivo.go com buffer em file

	// Verifica ocorrencia de erro
	if ferror != nil {
		fmt.Println("Erro:", ferror)
		panic(1)
	} else {
		reader := bufio.NewReader(file) // objeto reader para o buffer file
		file2, ferror := os.Create("arquivo2.txt")

		if ferror != nil {
			fmt.Println("Erro:", ferror)
			panic(2)
		} else {
			for {
				line, rerror := reader.ReadString('\n')

				fmt.Print(line)
				file2.WriteString(line) // escrevendo no arquivo2.txt

				if rerror == io.EOF {
					fmt.Println("FIM DO ARQUIVO")
					break
				}
			}
			if err := file2.Close(); err != nil {
				fmt.Println("Erro ao fechar arquivo:", err)
				panic(3)
			}
		}
	}

	// if <atributo> <expressao>
	if err := file.Close(); err != nil {
		fmt.Println("Erro ao fechar arquivo:", err)
		panic(4)
	}
}
