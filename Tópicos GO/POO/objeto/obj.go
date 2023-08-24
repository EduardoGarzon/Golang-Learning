// nao existem classes em go
// funcoes sao extendidas como metodos associados a uma struct
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// definicao de uma struct
type User struct {
	nome     string
	cadastro int
}

// metodos setters
// func (objeto *tipoObjeto) nomeFuncao(parametros tipoParametros) (retornos)
func (u *User) setNome(std string) {
	u.nome = std
}

func (u *User) setCadastro(cad int) {
	u.cadastro = cad
}

// metodos getters
// func (objeto tipoObjeto) nomeFuncao(parametros tipoParametros) (retornos)
func (u User) getNome() string {
	return u.nome
}

func (u User) getCadastro() int {
	return u.cadastro
}

func Clear() {
	fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}

func main() {
	var (
		nome     string
		cadastro int
	)

	// instancia do objeto
	p1 := User{nome, cadastro}

	// utilizacao dos metodos
	// objeto.nomeFuncao(parametros)

	fmt.Print("Informe o nome: ")
	scanner := bufio.NewReader(os.Stdin)
	line, _ := scanner.ReadString('\n')
	nome = strings.Trim(line, "\n")
	p1.setNome(nome) // atribui nome

	fmt.Print("Informe a matricula: ")
	fmt.Scan(&cadastro)
	p1.setCadastro(cadastro) // atribui cadastro

	fmt.Printf("NOME: %s\n", p1.getNome())
	fmt.Printf("CADASTRO: %d\n", p1.getCadastro())
}
