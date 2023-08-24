// heranca é implementada atraves de composicao e nao da forma tradicional
// defini-se uma estrutura pai e uma estrutura filha que contem a estrutura pai herdando seus metodos
package main

import (
	"fmt"
)

// estrutura pai
type Animal struct {
	sound string
}

// metodo da classe pai
func (a Animal) makeSound() string {
	a.sound = "ATUMALACA!"
	return a.sound
}

// estrutura filha
type Dog struct {
	Animal // inclui a estrutura pai herdando seus metodos
}

// metodo da classe filha
func (d Dog) Bark() string {
	bark := "AU AU AU!"
	return bark
}

func main() {
	// instanciando objeto
	animal := Dog{}

	fmt.Println(animal.makeSound())
	fmt.Println(animal.Bark())
}
