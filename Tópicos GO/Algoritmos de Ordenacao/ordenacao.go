/***********************************************************
 *Implemantação de algoritmos de ordenação genéricos em Go *
 ***********************************************************/
package main

import (
	"fmt"
	"reflect"
)

// Declara uma interface para a qual os algoritmos serão definidos.
// Qualquer classe que implementar as funções desta interface poderão
// usar estes algoritmos.
type Sorter interface {
	Len() int           // Função que retorna o comprimento do slice de dados a ser ordenado
	Less(i, j int) bool // Função de comparação a ser usada pelos algoritmos
	Swap(i, j int)      // Função de troca a ser usada pelos algoritmos
}

// Definição do slice de inteiros a ser usado como exemplo
type IntSlice []int

// Definição da função Len() para o slice de inteiros
func (intsl IntSlice) Len() int {
	return len(intsl)
}

// Definição da função Less() para o slice de inteiros
func (intsl IntSlice) Less(i, j int) bool {
	return intsl[i] < intsl[j]
}

// Definição da função Swap para o slice de inteiros
func (intsl IntSlice) Swap(i, j int) {
	intsl[i], intsl[j] = intsl[j], intsl[i]
}

// definição da classe pessoa
type Pessoa struct {
	codigo uint8
	nome   string
}

// Definição de um slice do tipo pessoa
type PessoaSlice []Pessoa

// Definição da função Len() para o slice Pessoa
func (ps PessoaSlice) Len() int {
	return len(ps)
}

// definição da função Less() para o slice Pessoa
func (ps PessoaSlice) Less(i, j int) bool {
	return ps[i].nome < ps[j].nome
}

// Definição da função Swap para o slice Pessoa
func (ps PessoaSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// Algoritmo de Ordenação bubbleSort (Método da Bolha)
func bubbleSort(s Sorter) {
	n := s.Len()
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			if s.Less(j, i) {
				s.Swap(i, j)
			}
		}
	}
}

// Algoritmo de Ordenação shakerSort
func shakerSort(s Sorter) {
	left := 0
	right := s.Len() - 1
	for left < right {
		for i := left; i < right; i++ {
			if s.Less(i + 1, i) {
				s.Swap(i, i + 1)
			}
		}
		right--

		for i := right; i > left; i-- {
			if s.Less(i, i - 1) {
				s.Swap(i, i - 1)
			}
		}
		left++
	}
}

// Algoritmo de ordenação selectionSort
func selectionSort(s Sorter) {
	n := s.Len()
	for i := 0; i < n - 1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if s.Less(j, minIndex) {
				minIndex = j
			}
		}
		if minIndex != i {
			s.Swap(i, minIndex)
		}
	}
}

// Algoritmo de ordenação insertionSort
func insertionSort(s Sorter) {
	n := s.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s.Less(j, j - 1); j-- {
			s.Swap(j, j - 1)
		}
	}
}

// Algoritmo de ordenação shellSort
func shellSort(s Sorter) {
	n := s.Len()
	gap := 1
	for gap < n/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < n; i++ {
			for j := i; j >= gap && s.Less(j, j-gap); j -= gap {
				s.Swap(j, j-gap)
			}
		}
		gap /= 3
	}
}

// Função shell (concha) para o MergeSort()
// Recebe um slice tipo interface vazia e chama o mergeSort passando os valores do slice, a posição inicial (0)
// e a posição final (.Len() - 1
/*
func MergeSort(slice Sorter) Sorter {
    if slice.Len() < 2 {
        return slice
    }
    mid := slice.Len() / 2
    left := MergeSort(slice.Slice(0, mid))
    right := MergeSort(slice.Slice(mid, slice.Len()))
    return Merge(left, right)
}

func Merge(left, right Sorter) Sorter {
    result := make(Pessoa, left.Len()+right.Len())

    i, j, k := 0, 0, 0
    for i < left.Len() && j < right.Len() {
        if left.Less(i, j) {
            result.Swap(k, i)
            i++
        } else {
            result.Swap(k, j)
            j++
        }
        k++
    }
	for i < left.Len() {
        result.Swap(k, i)
        i++
        k++
    }
    for j < right.Len() {
        result.Swap(k, j)
        j++
        k++
    }
    return result
}
*/

func SequentialSearch(slice Sorter, value interface{}) int {
    s := reflect.ValueOf(slice)
    if s.Kind() != reflect.Slice {
        panic("SequencialSearch: o primeiro argumento deve ser uma slice")
    }
    for i := 0; i < s.Len(); i++ {
        if reflect.DeepEqual(s.Index(i).Interface(), value) {
            return i
        }
   }
    return -1
}

func BinarySearch(slice Sorter, value interface{}) int {
    s := reflect.ValueOf(slice)
    if s.Kind() != reflect.Slice {
        panic("BinarySearch: o primeiro argumento deve ser uma slice")
    }

    low := 0
    high := s.Len() - 1
    for low <= high {
        mid := (low + high)/2
        if reflect.DeepEqual(s.Index(mid).Interface(), value) {
            return mid
        }
        if sorter, ok := slice.(Sorter); ok {
            if sorter.Less(mid, s.Len()) {
                high = mid - 1
            } else {
                low = mid + 1
            }
        }
    }
    return -1
}

func main() {
	// Exemplo de uso com slice de inteiros
	intSlice := IntSlice{100, 1, 99, 2, 98, 3, 97, 4, 96, 5, 95, 6, 94, 7, 93}
	bubbleSort(intSlice)
	fmt.Println("Posição do número 4 >>", SequentialSearch(intSlice, 4))
	//fmt.Println("Posição do número 4 >>", BinarySearch(intSlice, 4))
	fmt.Println(intSlice)

	// Exemplo de uso com slice de pessoas
	pessoaSlice := PessoaSlice{
		Pessoa{codigo: 1, nome: "zora"},
		Pessoa{codigo: 2, nome: "yara"},
		Pessoa{codigo: 3, nome: "xenia"},
		Pessoa{codigo: 4, nome: "walter"},
	}
	bubbleSort(pessoaSlice)
	fmt.Println(pessoaSlice)
	//a := BinarySearch(pessoaSlice, 1)
	//fmt.Println("Posição de walter >>", a)
	fmt.Println(intSlice)
}
