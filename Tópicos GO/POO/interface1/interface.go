// interfaces sao conjuntos de metodos que uma estrutura deve implementar
package main

import "fmt"

// interface
type animalFeeder interface {
	Feed()
}

// metodo da interface
// recebe um animal que implementa a interface animalFeeder
func feedAnimal(animal animalFeeder) {
	animal.Feed()
}

// estrutura
type Cat struct {
}

// estrutura cat agora contem a interface
func (c Cat) Feed() {
	fmt.Println("Animal Alimentado")
}

func main() {
	cat := Cat{}
	feedAnimal(cat) // como Car implementa a interface, pode utilizar os metodos da interface
}
