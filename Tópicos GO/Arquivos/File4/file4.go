package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var std string

	// abre arquivo
	file, ferror := os.Open("arquivo4.txt")
	if ferror != nil {
		fmt.Println("Erro ao abrir arquivo.")
		panic(ferror)
	}

	// fecha arquivo
	defer file.Close()

	// escreve no arquivo
	fmt.Fprintf(file, "Luiz")

	// sincroniza o conteudo do arquivo
	file.Sync()
	time.Sleep(time.Millisecond * 100)

	// le do arquivo
	fmt.Fscan(file, &std)
	fmt.Println(std)
}
