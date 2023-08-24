package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Aluno struct {
	nome      string
	matricula int8
}

func Clear() {
	fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}

func textColor(str string) {
	fmt.Print("\033[36m" + str + "\033[0m") // ciano text color
}

func Sscan(str *string) {
	scanner := bufio.NewReader(os.Stdin)
	line, _ := scanner.ReadString('\n')
	*str = strings.Trim(line, "\n")
}

func main() {
	var (
		alunos [10]Aluno
		total  int8
	)

	Clear()
	textColor("Informe o total de alunos (max. 10): ")
	fmt.Scanf("%d%*c", &total)

	for i := 0; i < int(total); i++ {
		Clear()
		textColor("Nome do aluno [" + strconv.Itoa(i+1) + "]: ")
		Sscan(&alunos[i].nome)
		textColor("Matricula: ")
		fmt.Scanf("%d%*c", &alunos[i].matricula)
	}

	Clear()
	fmt.Println("TURMA:")
	for i := 0; i < int(total); i++ {

		fmt.Printf("Nome: %s\n", alunos[i].nome)
		fmt.Printf("Matricula: %d\n", alunos[i].matricula)
	}
}
