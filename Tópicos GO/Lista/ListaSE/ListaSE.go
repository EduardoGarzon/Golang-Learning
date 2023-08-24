package main

import (
	"fmt"
)

/******************************************
 * Declaração da classe Node              *
 ******************************************/
type Node struct {
	data interface{}
	next *Node
}

/********************************
 * Declaração da classe ListaSE *
 ********************************/
type ListaSE struct {
	len   int8  // Número de elementos na lista
	left  *Node // ponteiro para o inicio da lista
	right *Node // ponteiro para o final da lista
}

/********************************
 * Verificação de lista vazia   *
 ********************************/
func (l *ListaSE) IsEmpty() bool {
	return l.len == 0
}

/*****************************************
 * Verifica se um dado é membro da lista *
 *****************************************/
func (l *ListaSE) IsMember(x interface{}) *Node {
	aux := l.left
	for (aux != nil) && aux.data == x {
		aux = aux.next
	}
	return aux
}

/******************************
 * Retorna o tamanho da lista *
 ******************************/
func (l *ListaSE) Length() int8 {
	return l.len
}

/*******************
 * Imprime a lista *
 *******************/
func (l *ListaSE) Print() {
	p := l.left
	fmt.Print("{")
	for p != nil {
		fmt.Printf("%v", p.data)
		if p.next != nil {
			fmt.Print(", ")
		}
		p = p.next
	}
	fmt.Println("}")
}

/***********************
 * Inserção à esquerda *
 ***********************/
func (list *ListaSE) InsertLeft(x interface{}) {
	newNode := new(Node)
	if newNode == nil {
		panic("Não há memória suficiente para alocar um novo nó!")
	}
	newNode.data = x

	// Se a lista estiver vazia, o novo nó será o primeiro e o último nó.
	if list.left == nil {
		newNode.next = nil
		list.left = newNode
		list.right = newNode
	} else {
		// Se a lista não estiver vazia, adicione o novo nó ao início da lista.
		newNode.next = list.left
		list.left = newNode
	}
	// incrementar o contador de elementos na lista
	list.len++
}

/***********************
 * Inserção à direita  *
 ***********************/
func (list *ListaSE) InsertRight(x interface{}) {
	newNode := new(Node) // alocar memória para um novo nó
	if newNode == nil {  // verifica se ocorreu erro de memória insuficiente
		panic("Não há memória suficiente para alocar um novo nó!")
	}
	// inicializar o valor e o próximo ponteiro do novo nó
	newNode.data = x
	newNode.next = nil

	// se a lista está vazia, o novo nó será a cabeça e a cauda da lista
	if list.left == nil {
		list.left = newNode
		list.right = newNode
	} else {
		// se a lista não está vazia, adicionar o novo nó no final da lista
		list.right.next = newNode
		list.right = newNode
	}

	// incrementar o contador de elementos na lista
	list.len++
}

/***********************
 * Remoção à esquerda  *
 ***********************/
func (l *ListaSE) RemoveLeft() interface{} {
	if l.left == nil {
		return nil
	}
	r := l.left.data
	l.left = l.left.next
	if l.left == nil {
		l.right = nil
	}
	l.len--
	return r
}

/***********************
 * Remoção à direita   *
 ***********************/
func (l *ListaSE) RemoveRight() interface{} {
	if l.right == nil {
		return nil
	}
	r := l.right.data
	if l.left == l.right {
		l.left = nil
		l.right = nil
	} else {
		aux := l.left
		for aux.next != l.right {
			aux = aux.next
		}
		aux.next = nil
		l.right = aux
	}
	l.len--
	return r
}

func main() {
	list := ListaSE{}

	// varifica se a lista está vazia
	fmt.Println("A lista está vazia? >>", list.IsEmpty())

	// verifica se 0 faz parte da lista
	fmt.Println("0 faz parte da lista? >>", list.IsMember(0))

	// verifica o comprimento da lista
	fmt.Println("Quantos elementos tem a lista? >>", list.Length())

	// imprimindo a lista
	fmt.Println("Imprimindo a lista >>")
	list.Print()
	fmt.Println()

	// inserindo à esquerda
	list.InsertLeft(0)
	list.InsertLeft(1)
	list.InsertLeft(2)

	// varifica se a lista está vazia
	fmt.Println("A lista está vazia? >>", list.IsEmpty())

	// verifica se 0 faz parte da lista
	fmt.Println("0 faz parte da lista? >>", list.IsMember(0))

	// verifica o comprimento da lista
	fmt.Println("Quantos elementos tem a lista? >>", list.Length())

	// imprimindo a lista
	fmt.Println("Imprimindo a lista >>")
	list.Print()
	fmt.Println()

	// inserindo à direita
	list.InsertRight(3)
	list.InsertRight(4)
	list.InsertRight(5)

	// varifica se a lista está vazia
	fmt.Println("A lista está vazia? >>", list.IsEmpty())

	// verifica se 0 faz parte da lista
	fmt.Println("0 faz parte da lista? >>", list.IsMember(0))

	// verifica o comprimento da lista
	fmt.Println("Quantos elementos tem a lista? >>", list.Length())

	// imprimindo a lista
	fmt.Println("Imprimindo a lista >>")
	list.Print()
	fmt.Println()

	// Removendo à esquerda
	r := list.RemoveLeft()
	if r != nil {
		fmt.Println("Elemento removido:", r)
	}

	// varifica se a lista está vazia
	fmt.Println("A lista está vazia? >>", list.IsEmpty())

	// verifica se 0 faz parte da lista
	fmt.Println("0 faz parte da lista? >>", list.IsMember(0))

	// verifica o comprimento da lista
	fmt.Println("Quantos elementos tem a lista? >>", list.Length())

	// imprimindo a lista
	fmt.Println("Imprimindo a lista >>")
	list.Print()
	fmt.Println()

	// Removendo à direita
	r = list.RemoveRight()
	if r != nil {
		fmt.Println("Elemento removido:", r)
	}

	// varifica se a lista está vazia
	fmt.Println("A lista está vazia? >>", list.IsEmpty())

	// verifica se 0 faz parte da lista
	fmt.Println("0 faz parte da lista? >>", list.IsMember(0))

	// verifica o comprimento da lista
	fmt.Println("Quantos elementos tem a lista? >>", list.Length())

	// imprimindo a lista
	fmt.Println("Imprimindo a lista >>")
	list.Print()
	fmt.Println()
}
