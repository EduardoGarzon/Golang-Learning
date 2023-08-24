package main

import (
	"fmt"
)

func main() {
	SIZE := 5
	q := make([]interface{}, SIZE)

	// Adicionar elementos à fila
	q[0] = 10
	q[1] = "Hello"
	q[2] = true

	// Acessar elementos da fila
	element1 := q[0].(int)
	element2 := q[1].(string)
	element3 := q[2].(bool)

	fmt.Println(element1, element2, element3)
}
