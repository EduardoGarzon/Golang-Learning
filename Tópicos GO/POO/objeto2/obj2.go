package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// declarando um tipo personalizado
type Car struct {
	brand string
	model string
	year  int
	age   int
}

// metodos setters
// func (objeto *tipoObjeto) nomeFuncao(parametros tipoParametros) (retornos)
func (c *Car) setBrand(std string) {
	c.brand = std
}

func (c *Car) setModel(std string) {
	c.model = std
}

func (c *Car) setYear(ano int) {
	c.year = ano
}

func (c *Car) setAge(ano int) {
	c.age = time.Now().Year() - ano
}

// metodos getters
// func (objeto tipoObjeto) nomeFuncao(parametros tipoParametros) (retornos)
func (c Car) getBrand() string {
	return c.brand
}

func (c Car) getModel() string {
	return c.model
}

func (c Car) getYear() int {
	return c.year
}

func (c Car) getAge() int {
	return time.Now().Year() - c.getYear()
}

// funcoes main
func cadastrarMarca(myCar *Car) {
	var marca string
	fmt.Print("Informe a marca do carro: ")
	scanner := bufio.NewReader(os.Stdin)
	line, _ := scanner.ReadString('\n')
	marca = strings.Trim(line, "\n")
	myCar.setBrand(marca)

}

func cadastrarModelo(myCar *Car) {
	var modelo string
	fmt.Print("Informe o modelo do carro: ")
	scanner := bufio.NewReader(os.Stdin)
	line, _ := scanner.ReadString('\n')
	modelo = strings.Trim(line, "\n")
	myCar.setModel(modelo)
}

func cadastrarAno(myCar *Car) {
	var ano int
	fmt.Print("Informe o ano do carro: ")
	fmt.Scan(&ano)
	myCar.setYear(ano)
	myCar.setAge(ano)
}

func showData(myCar *Car) {
	Clear()
	fmt.Println("MARCA:", myCar.getBrand())
	fmt.Println("MODELO:", myCar.getModel())
	fmt.Println("ANO:", myCar.getYear())
	fmt.Println("IDADE:", myCar.getAge())
}

func Clear() {
	fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}

func main() {
	// instanciando objeto
	myCar := Car{}

	cadastrarMarca(&myCar)
	cadastrarModelo(&myCar)
	cadastrarAno(&myCar)
	showData(&myCar)
}
