/***********************************************************
 * StsticStack.go                                          *
 * Implementação de uma pilha com vetor estático em Golang *
 ***********************************************************/
package main

import (
	"fmt"
)

/***********************************
 * Definição da classe StaticStack *
 ***********************************/
const SIZE = 10 // definição do tamanho da pilha

type StaticStack struct {
	s[SIZE] interface{} 		// estrutura que armazena a fila
	topo byte					// índices para topo da pilha
	tamanho	byte				// altura da pilha
}

// Inicialização da pilha
func (ss *StaticStack) InitStack() {
	ss.topo, ss.tamanho = 0, 0
}

// Impressão da pilha
func (ss *StaticStack) Println() {
	// A pilha sempre será impressa com a base à esquerda e o topo à direita
	fmt.Print("{")
	var l byte = 0;
	for k := byte(0); k < ss.topo; k++ {
		fmt.Print(ss.s[k])
		if l < (ss.tamanho - 1)  {
			fmt.Print(",")
		}
		l++
	}
	fmt.Println("}")
}

// Pilha vazia
func (ss *StaticStack) IsEmpty() bool {
	return ss.topo == 0
}

// Pilha cheia
func (ss *StaticStack) IsFull() bool {
	return ss.topo == SIZE
}

// Push - empilhar
func (ss *StaticStack) Push(x interface{}) {
	if ss.IsFull() {
		panic("Erro: Pilha cheia");
	} else {
		ss.s[ss.topo] = x
		ss.topo++
		ss.tamanho++
	}
}

// pop - desempilhar
func (ss *StaticStack) Pop() (x interface{}) {
	if ss.IsEmpty() {
		panic("Erro: Pilha vazia!")
	} else {
		ss.topo--
		x = ss.s[ss.topo]
		ss.tamanho--
		return x
	}
}

/*****************************
 * Fim da classe StaticStack *
 *****************************/
 
func main() {
	var pilha StaticStack
	
	// inicialização da pilha
	pilha.InitStack()
	
	// Imprime a pilha inicial
	fmt.Print("Pilha = ")
	pilha.Println()
	
	// Empilha 0
	fmt.Println("Empilhando 0:")
	pilha.Push(0)
	pilha.Println()
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())
	
	// Empilha 1
	fmt.Println("Empilhando 1:")
	pilha.Push(1)
	pilha.Println()
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())

	// Empilha 2
	fmt.Println("Empilhando 2:")
	pilha.Push(2)
	pilha.Println()
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())

	// Empilha 3
	fmt.Println("Empilhando 3:")
	pilha.Push(3)
	pilha.Println()
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())

	// Empilha 4
	fmt.Println("Empilhando 3:")
	pilha.Push(4)
	pilha.Println()
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())

	// Empilha 5
	fmt.Println("Empilhando 3:")
	pilha.Push(5)
	pilha.Println()
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())

	// Desempilha
	fmt.Println("Desempilhando:")
	b := pilha.Pop()
	pilha.Println()
	fmt.Println("Removido:", b)
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())

	// Empilha 6
	fmt.Println("Empilhando 6:")
	pilha.Push(6)
	pilha.Println()
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())

	// Desempilha
	fmt.Println("Desempilhando:")
	b = pilha.Pop()
	pilha.Println()
	fmt.Println("Removido:", b)
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())

	// Desempilha
	fmt.Println("Desempilhando:")
	b = pilha.Pop()
	pilha.Println()
	fmt.Println("Removido:", b)
	fmt.Println("Pilha vazia? ", pilha.IsEmpty())
	fmt.Println("Pilha cheia? ", pilha.IsFull())
}	 
