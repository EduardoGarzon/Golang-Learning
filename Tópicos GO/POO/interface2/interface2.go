package main

import (
	"fmt"
	"math"
)

// interface
type Shape interface {
	Area() float64
}

// estrutura 1
type Rectangle struct {
	Width  float64
	Height float64
}

// estrutura 1 implementa metodo Area() da interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// estrutura 2
type Circle struct {
	Radius float64
}

// estrutura 2 implementa o metodo Area() da interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// funcao da Interface para printar o metodo Area()
func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	// instanciando structs 1 e 2
	rect := Rectangle{Width: 2, Height: 4}
	circle := Circle{Radius: 3}

	// ambas as structs contem a interface
	PrintArea(rect)   // Saída: Area: 8
	PrintArea(circle) // Saída: Area: 28.274333882308138

	// rect   -> Interface (Area()) -> func (r Rectangle) Area()
	// circle -> Interface (Area()) -> func (c Circle) Area()

}
