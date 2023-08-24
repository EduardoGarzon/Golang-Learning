package main

import "fmt"

func main() {

	collection := []string{"Luiz", "Eduarda", "Heloisa", "Pedro"}
	collection2 := []int{20, 21, 20, 19}

	// range percorre uma estrutura de dados
	// retorna o indice e o valor do elemento
	for index, value := range collection {
		fmt.Println(index, value)
	}

	// podemos omitir os valores que queremos
	for _, value := range collection2 {
		fmt.Println(value)
	}

	fmt.Println(collection)
	fmt.Println(collection2)
}
