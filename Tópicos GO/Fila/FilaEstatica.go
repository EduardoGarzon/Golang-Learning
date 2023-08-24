/**************************************************
 * StaticQueue.go                                 *
 * Implementa uma fila baseada em vetor em Golang *
 **************************************************/

package main

import (
	"fmt"
)

/***********************************
 * Definição da classe StaticQueue *
 ***********************************/
const SIZE = 10 // definição do tamanho da fila

type StaticQueue struct {
	q           [SIZE]interface{} // estrutura que armazena a fila
	inicio, fim byte              // índices para o início e o fim da fila
	tamanho     byte              // comprimento da fila
}

// Inicialização da fila
func (sq *StaticQueue) InitQueue() {
	sq.inicio, sq.fim, sq.tamanho = 0, 0, 0
}

// Impressão da fila
func (sq *StaticQueue) Println() {
	fmt.Print("{")
	var l byte = 0
	for k := sq.inicio; k != sq.fim; k = (k + 1) % SIZE {
		fmt.Print(sq.q[k])
		if l < (sq.tamanho - 1) {
			fmt.Print(",")
		}
		l++
	}
	fmt.Println("}")
}

// Verifica fila vazia
func (sq *StaticQueue) IsEmpty() bool {
	return (sq.inicio == sq.fim)
}

// verifica Fila cheia
func (sq *StaticQueue) IsFull() bool {
	return (sq.fim+1)%SIZE == sq.inicio
}

// Insere
func (sq *StaticQueue) Insert(x interface{}) {
	if sq.IsFull() {
		panic("Erro: Fila cheia!")
	} else {
		sq.q[sq.fim] = x
		sq.fim = (sq.fim + 1) % SIZE
		sq.tamanho++
		return
	}
}

// remove
func (sq *StaticQueue) Remove() (r interface{}) {
	if sq.IsEmpty() {
		panic("Erro: Fila vazia!")
	} else {
		r = sq.q[sq.inicio]
		sq.inicio++
		sq.tamanho--
		return r
	}
}

/*****************************
 * Fim da classe StaticQueue *
 *****************************/

// Função principal
func main() {
	var fila StaticQueue // variável fila do tipo struct StaticQueue

	// Inicialização da fila
	fila.InitQueue()

	// Imprime a fila inicial
	fmt.Print("Fila = ")
	fila.Println()

	// Fila está vazia?
	fmt.Println("Fila vazia? ", fila.IsEmpty())

	// Fila está cheia?
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 0:")
	fila.Insert(0)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 1:")
	fila.Insert(1)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 2:")
	fila.Insert(2)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 3:")
	fila.Insert(3)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 4:")
	fila.Insert(4)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 5:")
	fila.Insert(5)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 6:")
	fila.Insert(6)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 7:")
	fila.Insert(7)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 8:")
	fila.Insert(8)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// remove da fila
	fmt.Println("removendo da fila:")
	a := fila.Remove()
	fmt.Println(a, "removido da fila!")
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// remove da fila
	fmt.Println("removendo da fila:")
	a = fila.Remove()
	fmt.Println(a, "removido da fila!")
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 9:")
	fila.Insert(9)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 10:")
	fila.Insert(10)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// remove da fila
	fmt.Println("removendo da fila:")
	a = fila.Remove()
	fmt.Println(a, "removido da fila!")
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// inserindo na fila
	fmt.Println("Inserindo 11:")
	fila.Insert(11)
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// remove da fila
	fmt.Println("removendo da fila:")
	a = fila.Remove()
	fmt.Println(a, "removido da fila!")
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())

	fmt.Println()

	// remove da fila
	fmt.Println("removendo da fila:")
	a = fila.Remove()
	fmt.Println(a, "removido da fila!")
	fmt.Print("Fila = ")
	fila.Println()
	fmt.Println("Fila vazia? ", fila.IsEmpty())
	fmt.Println("Fila cheia? ", fila.IsFull())
}
