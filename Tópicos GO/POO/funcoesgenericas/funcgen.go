package main

import "fmt"

// funcao que verifica se determinado valor esta no slice
// funcao generica com paremetro de tipo T para o slice e o value
// T é o tipo comparavel para operacao de comparacao
// assim a funcao pode receber qualquer valor generico como int, float e strings
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Contains(slice1, 3)) // Saída: true

	slice2 := []string{"foo", "bar", "baz"}
	fmt.Println(Contains(slice2, "qux")) // Saída: false
}
