package main

import "fmt"

func main() {
	var (
		A []byte // slice vazio
		B = []int{1, 2, 3, 4, 5}
	)
	C := []float64{1.5, 2.5, 3.5, 4.5, 5.5} // declaracao curta

	// inicializando o slice A
	for i := 0; i < 5; i++ {
		//A[i] = byte(i) NAO FUNCIONA
		// utilizamos append para acrescimoes em slices nao inicializados
		A = append(A, byte(i))
	}

	fmt.Println("Slice A:", A)
	fmt.Println("Slice B:", B)
	fmt.Println("Slice C:", C)

	//Atualizando os slices
	for i := 0; i < len(A); i++ {
		A[i] = byte(i + 1)
		B[i] = int(i)
		C[i] = float64(i)
	}

	fmt.Println("Slice A:", A)
	fmt.Println("Slice B:", B)
	fmt.Println("Slice C:", C)
}
