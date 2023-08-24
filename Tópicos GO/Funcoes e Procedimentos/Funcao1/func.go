package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Clear() {
	fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}

func takeInfo() (string, int) {
	Clear()
	var nome string
	var quantidade int
	fmt.Print("Informe seu nome: ")
	scanner := bufio.NewReader(os.Stdin)
	line, _ := scanner.ReadString('\n')
	nome = strings.Trim(line, "\n")
	fmt.Print("Informe a quantidade de notas: ")
	fmt.Scan(&quantidade)
	return nome, quantidade
}

func takeGrades(slice *[]float64, qtd *int) {
	Clear()
	var grade float64
	for i := 0; i < *qtd; i++ {
		fmt.Print("Informe a nota [" + strconv.Itoa(i+1) + "]: ")
		fmt.Scan(&grade)
		*slice = append(*slice, grade)
	}
}

func makeMedia(slice *[]float64, qtd *int) float64 {
	var media float64
	var soma float64
	for i := 0; i < *qtd; i++ {
		soma += (*slice)[i]
	}
	media = soma / float64(*qtd)
	return media
}

func showData(media *float64, grades *[]float64, nome *string) {
	Clear()
	fmt.Println("Nome: ", *nome)
	fmt.Println("Notas: ", *grades)
	fmt.Println("Media: ", *media)
}

func main() {
	var (
		nome       string
		quantidade int
		grades     = make([]float64, 0, quantidade)
		media      float64
	)

	nome, quantidade = takeInfo()
	takeGrades(&grades, &quantidade)
	media = makeMedia(&grades, &quantidade)
	showData(&media, &grades, &nome)

}
