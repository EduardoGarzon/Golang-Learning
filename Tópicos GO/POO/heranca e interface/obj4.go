package main

import (
	"fmt"
	"os"
	"time"
)

// interface
type animalCare interface {
	Feed()
}

// metodos da interface
func feedAnimal(animal animalCare) {
	animal.Feed()
}

// estrutura pai
type Animal struct {
	energia int
}

// Animal implementa interface
func (a Animal) Feed() string {
	feed := "Animal Alimentado!"
	return feed
}

// estrutura filha
type Dog struct {
	Animal // herda Animal
}

// metodos da estrutura pai
func (a Animal) Jump() string {
	jump := "JUMPED!"
	return jump
}

// metodos da estrutura filha
func makeDog(dog *Dog) Dog {
	*dog = Dog{}
	*&dog.energia = 10
	return *dog
}

func (d Dog) Bark() string {
	bark := "AU AU AU!"
	return bark
}

func (d Dog) dogStatus() (string, int) {
	var msg string
	var e int = d.energia
	if d.energia >= 5 {
		msg = fmt.Sprintf("NOT HUNGRY :) (ENERGIA: %d)", e)
	} else if d.energia < 5 {
		msg = fmt.Sprintf("HUNGRY :) (ENERGIA: %d)", e)
	}
	return msg, e
}

func (d Dog) getEnergia() int {
	return d.energia
}

func (d *Dog) setEnergia(food int) {
	d.energia += food
}

func Menu() {
	var option int

	Clear()
	fmt.Println("\033[36m--- DOG SIMULATOR ---")
	fmt.Println("[1] BARK")
	fmt.Println("[2] JUMP")
	fmt.Println("[3] FEED")
	fmt.Println("[4] DOG STATUS")
	fmt.Println("[5] EXIT")
	fmt.Print("Choose option: \033[0m")
	fmt.Scanln(&option)
	Clear()

	switch option {
	case 1:
		if dog.getEnergia() >= 5 {
			fmt.Println(dog.Bark())
			dog.setEnergia(-1)
		} else {
			fmt.Println(dog.dogStatus())
		}
		fmt.Print("Tecle [ENTER]")
		fmt.Scanln()
		Menu()
	case 2:
		if dog.getEnergia() >= 5 {
			fmt.Println(dog.Jump())
			dog.setEnergia(-2)
		} else {
			fmt.Println(dog.dogStatus())
		}
		fmt.Print("Tecle [ENTER]")
		fmt.Scanln()
		Menu()
	case 3:
		if dog.getEnergia() >= 5 {
			dog.dogStatus()
		} else {
			fmt.Println(dog.dogStatus())
			for dog.getEnergia() < 10 {
				dog.setEnergia(1)
				fmt.Println("ENERGIA:", dog.getEnergia())
			}
			fmt.Println(dog.Feed())
			fmt.Println(dog.dogStatus())
		}
		fmt.Print("Tecle [ENTER]")
		fmt.Scanln()
		Menu()
	case 4:
		msg, _ := dog.dogStatus()
		fmt.Println(msg)
		fmt.Print("Tecle [ENTER]")
		fmt.Scanln()
		Menu()
	case 5:
		fmt.Println("Encerrando simulador.")
		time.Sleep(time.Millisecond * 500)
		os.Exit(1)
	default:
		fmt.Println("Opcao Invalida, tente novamente!")
		fmt.Print("Tecle [ENTER]")
		fmt.Scanln()
		Menu()
	}
}

// limpa tela
func Clear() {
	fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}

var (
	dog Dog
)

func main() {
	dog = makeDog(&dog)
	Menu()
}
