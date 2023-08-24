/******************************************************************
 * BinTree.go: Implementação de uma árvore binária de busca em go *
 ******************************************************************/
package main

import(
   "fmt"
)

/*****************************************
 * Declaração do tipo de dado armazenado *
 *****************************************/
 type DataType uint8

/**************************************
 * Declaração de nó de árvore Binária *
 **************************************/
type BinTreeNode struct {
   data DataType // dado armazenado no nó
   left, right *BinTreeNode	// ponteiros para os filhos esquerdo e direito
}

/****************************************************
 * Declaração de nó de Lista simplesmente encadeada *
 ****************************************************/
type Node struct {
   data *BinTreeNode
   next *Node
}

/********************************************
 * Declaração dos métodos da classe ListaSE *
 ********************************************/
type ListaSE struct {
   len int8	// Número de elementos na lista
   left *Node	// ponteiro para o inicio da lista
   right *Node	// ponteiro para o final da lista
}

/******************************
 * Verificação de lista vazia *
 ******************************/
func (l *ListaSE) IsEmpty() bool {
   return l.len == 0
}

/*****************************************
 * Verifica se um dado é membro da lista *
 *****************************************/
func (l *ListaSE) Find(x *BinTreeNode) *Node {
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
func (list *ListaSE) InsertLeft(x *BinTreeNode) {
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
func (list *ListaSE) InsertRight(x *BinTreeNode) {
   newNode := new(Node) // alocar memória para um novo nó
   
   if newNode == nil { // verifica se ocorreu erro de memória insuficiente
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
func (l *ListaSE) RemoveLeft() *BinTreeNode {
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
func (l *ListaSE) RemoveRight() *BinTreeNode {
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

/*********************************************************
 * init(): Função de inicialização da classe BinTreeNode *
 *********************************************************/
func (btn *BinTreeNode)Init(l *BinTreeNode, d DataType, r *BinTreeNode) {
   btn.data = d
   btn.left = l
   btn.right = r
}

/**********************************************
 * GetData(): retorna o dado armazenado no nó *
 **********************************************/
func (btn *BinTreeNode)GetData()(data DataType) {
   return btn.data
}

/****************************************************
 * SetData(): seta o valor do dado armazenado no nó *
 ****************************************************/
func (btn *BinTreeNode)SetData(d DataType) {
   btn.data = d
}

/*************************************************************
 * GetLeft(): retorna o ponteiro para o filho esquerdo do nó *
 *************************************************************/
func (btn *BinTreeNode)GetLeft()(l *BinTreeNode) {
   return btn.left
}

/*************************************************************
 * GetRight(): retorna o ponteiro para o filho direito do nó *
 *************************************************************/
func (btn *BinTreeNode)GetRight()(r *BinTreeNode) {
   return btn.right
}

/******************************************************
 * SetLeft(): Seta o ponteiro do filho esquerdo do nó *
 ******************************************************/
func (btn *BinTreeNode)SetLeft(l *BinTreeNode) {
   btn.left = l
}

/******************************************************
 * SetRight(): Seta o ponteiro do filho direito do nó *
 ******************************************************/
func (btn *BinTreeNode)SetRight(r *BinTreeNode) {
   btn.right = r
}

/*************************************
 * Print(): Imprime o conteúdo do nó *
 *************************************/
func (btn *BinTreeNode)Print() {
   fmt.Print(btn.data)
}
 
/************************************************************
 * Implementação da classe BinTree: árvore binária de busca *
 ************************************************************/
type BinTree struct {
   root *BinTreeNode
   numElements int8
}

/*******************************************
 * Init(): Inicialização da classe BinTree *
 *******************************************/
func (bt *BinTree) Init()  {
   bt.root = nil
   bt.numElements = 0
}

/******************************************************
 * GetRoot(): Retorna o endereço do nó raiz da árvore *
 ******************************************************/
func (bt *BinTree) GetRoot() *BinTreeNode {
   return bt.root
}

/*************************************************************************
 * GetNumElements(): retorna o número de elementos armazenados na árvore *
 *************************************************************************/
func (bt *BinTree) GetNumElements() int8 {
   return bt.numElements
}

/***************************************
 * Print(): Imprime a árvore formatada *
 ***************************************/
func (bt *BinTree)Print(r *BinTreeNode, level int8) {
   if r == nil {
      return
   } else {
      bt.Print(r.GetRight(), level + 1)
      for i := int8(0); i < level; i++ {
	 fmt.Print("------|")
      }
      r.Print() // imprime o nó
      fmt.Print("\n")
      bt.Print(r.GetLeft(), level + 1) 
   }
}

/*****************************************************************************
 * insertNode(): método recursivo auxiliar para inserir um nó em sua posição * 
 * correta na árvore                                                         *
 *****************************************************************************/
func (bt *BinTree)InsertNode(x DataType, root *BinTreeNode)(*BinTreeNode) {
   if root == nil { // nó nulo, tenta alocar um novo nó
      root = new(BinTreeNode)
      if root == nil {
	 panic("Erro: Memória insuficiente!\n")
      } else {
	 root.Init(nil, x, nil)  // novo nó inserido
	 bt.numElements++			// incrementa número de elementos
	 return root					// retorna o novo nó
      }
   } else { // nó não é nulo, verifica se tem que descer à esquerda ou à direita
      if x < root.GetData() {
	 root.SetLeft(bt.InsertNode(x, root.GetLeft()))
      } else {
	 root.SetRight(bt.InsertNode(x, root.GetRight()))
      }
      return root
   }
}

/******************************************
 * Insert(): Método de inserção na árvore *
 ******************************************/
func (bt *BinTree)Insert(x DataType) {
   bt.root = bt.InsertNode(x, bt.root)
}

/***************************************************************************
 * remNode(): Método recursivo auxiliar para remover um elemento da árvore *
 ***************************************************************************/
func (bt *BinTree) remNode(x DataType, r *BinTreeNode) *BinTreeNode {
   var (
      p1, p2 *BinTreeNode
   )
	
   if r == nil { // raiz é nula. Elemento não está na árvore
      return nil // Nada a fazer.
   }
   if r.GetData() == x { // apaga a raiz
      if r.GetLeft() == r.GetRight() { // não tem filhos (ambos são nulos)
	 r = nil // remove o nó r
	 bt.numElements-- // decrementa o número de elementos
	 return nil	// retorna nó nulo
      } else {	// algum filho é nulo
	 if r.GetLeft() == nil { // esquerda nula
	    p1 = r.GetRight()
	    r = nil
	    bt.numElements--
	    return p1
	 } else {
	    if r.GetRight() == nil {
	       p1 = r.GetLeft()
	       r = nil
	       bt.numElements--
	       return p1
	    } else { // os dois filhos estão presentes
	       p1 = r.GetRight()
	       p2 = r.GetRight()
	       for p1.GetLeft() != nil {
		  p1 = p1.GetLeft()
	       }
	       p1.SetLeft(r.GetLeft())
	       r = nil
	       bt.numElements--
	       return p2
	    }
	 }
      }
   }
   if (x < r.GetData()) {
      r.SetLeft(bt.remNode(x, r.GetLeft()))
   } else {
      r.SetRight(bt.remNode(x, r.GetRight()))
   }
   return r
}

/*************************************************************************
 * Delete(): Função que remove o elemento x da árvore (apenas a primeira *
 * ocorrencia)                                                           *
 *************************************************************************/
func (bt *BinTree) Delete(x DataType) {
   bt.root = bt.remNode(x, bt.root)
}

/****************************************************************************
 * Find(): retorna o endereço do elemento na árvore, se o mesmo estiver * 
 * na árvore, ou nil caso o elemento não esteja na árvore                   *
 ****************************************************************************/
 func (bt *BinTree) Find(x DataType) *BinTreeNode {
   r := bt.root
   achou := false
	
   for (!achou) && (r != nil) {
      if r == nil {
	 return nil
      } else {
	 if r.GetData() == x { //achou
	    achou = true;
         } else {
            if r.GetData() > x { // busca na esquerda
	       r = r.GetLeft()
            } else {
	       r = r.GetRight();
            }
         }
      }
   }
   return r
}

/**************************
 * Algoritmos de percurso *
 **************************/
 
/***********************************************************************
 * Algoritmo de visitação de nós utilizado pelos algoritmos iterativos *
 * R é o nó que está sendo visitado.                                   *
 * Este método é usado em todos os algoritmos de percurso              *
 ***********************************************************************/
func (bt *BinTree) Visit(r *BinTreeNode) {
   // Implemente aqui o processo de visita ao nó.
   // Neste exemplo, a visita apenas imprime o conteúdo armazenado no nó.
   fmt.Print(r.GetData(), " ")
}

/*****************************************************************************
 * BreadthFirst(): Realiza um percurso em largura iterativo na árvore.       *
 * Utiliza uma fila para realizar este procedimento. Imprime os nós na ordem *
 * em que o percurso é realizado                                             *
 *****************************************************************************/
func (bt *BinTree) BreadthFirst() {
   var (
      r *BinTreeNode = bt.root
      fila  ListaSE
   )
		        
   if (r != nil) {
      fila.InsertRight(r)
      fmt.Print("{ ")
      for !fila.IsEmpty() {
	 r = fila.RemoveLeft()
         bt.Visit(r)
         if r.GetLeft() != nil {
	    fila.InsertRight(r.GetLeft())
         }
         if (r.GetRight() != nil) {
	    fila.InsertRight(r.GetRight())
         }
      }
      fmt.Print("}")
   }
   return
}

/****************************************
 * Percursos em profundidade iterativos *
 ****************************************/
 
/********************************
 * Percurso Pré-ordem iterativo *
 ********************************/
func (bt *BinTree)ItPreOrder() {
   var (
      r *BinTreeNode = bt.root
      pilha ListaSE
   )
        
   if (r != nil) {
      pilha.InsertLeft(r)
      fmt.Print("{ ")
      for !pilha.IsEmpty() {
	 r = pilha.RemoveLeft()
	 bt.Visit(r)
	 if (r.GetRight() != nil) {
	    pilha.InsertLeft(r.GetRight())
	 }
	 if (r.GetLeft() != nil) {
	    pilha.InsertLeft(r.GetLeft())
	 }
      }
      fmt.Print("}")
   }
   return
}

/*******************************
 * Percurso em-Ordem iterativo *
 *******************************/
func (bt *BinTree) ItInOrder() {
   var (
      r *BinTreeNode = bt.root
      pilha ListaSE
   )
   fmt.Print("{ ")  
   for r != nil {
      for r != nil {
	 if (r.GetRight() != nil) {
	    pilha.InsertLeft(r.GetRight())
         }
         pilha.InsertLeft(r)
         r = r.GetLeft()
      }
      r = pilha.RemoveLeft()
      for !pilha.IsEmpty() && (r.GetRight() == nil) {
	 bt.Visit(r)
	 r = pilha.RemoveLeft()
      }
      bt.Visit(r)
      if !pilha.IsEmpty() {
	 r = pilha.RemoveLeft()
      } else {
	 r = nil
      }
   }
   fmt.Print("}")
   return
}

/********************************
 * Percurso Pós-Ordem Iterativo *
 ********************************/
func (bt *BinTree) ItPostOrder() {
   var (
      pilha ListaSE
      r *BinTreeNode = bt.root
      rAux *BinTreeNode = bt.root
   )
   
   fmt.Print("{ ")
   for r != nil {
      for r.GetLeft() != nil {
	 pilha.InsertLeft(r)
         r = r.GetLeft()
      }
      for (r != nil) && ((r.GetRight() == nil) || (r.GetRight() == rAux)) {
	 bt.Visit(r)
         rAux = r
         if pilha.IsEmpty() {
	    fmt.Print("}")
	    return
         }
         r = pilha.RemoveLeft()
      }
      pilha.InsertLeft(r)
      r = r.GetRight()
   }
}

/*****************************************************
 * Algoritmos de percurso em profundidade recursivos *
 *****************************************************/
 
/*************
 * Pré-Ordem *
 *************/
 
// Função recursiva interna
func (bt *BinTree) rPreOrd(r *BinTreeNode) {
   if (r == nil) {
      return
   } else {
      bt.Visit(r)
      bt.rPreOrd(r.GetLeft())
      bt.rPreOrd(r.GetRight())
   }
}

// Função que deve ser chamada
func (bt *BinTree)RecPreOrder() {
   r := bt.root
   fmt.Print("{ ")
   bt.rPreOrd(r)
   fmt.Print("}")
}

/************
 * Em-ordem *
 ************/
 
// Função recursiva interna
func (bt *BinTree) rInOrd(r *BinTreeNode) {
   if (r == nil) {
      return
   } else {
      bt.rInOrd(r.GetLeft())
      bt.Visit(r)
      bt.rInOrd(r.GetRight())
   }
}

// Função que deve ser chamada
func (bt *BinTree)RecInOrder() {
   r := bt.root
   fmt.Print("{ ")
   bt.rInOrd(r)
   fmt.Print("}")
}

/*************
 * Pós-ordem *
 *************/

// Função recursiva interna
func (bt *BinTree) rPosOrd(r *BinTreeNode) {
   if (r == nil) {
      return
   } else {
      bt.rPosOrd(r.GetLeft())
      bt.rPosOrd(r.GetRight())
      bt.Visit(r)
   }
}

// Função que deve ser chamada
func (bt *BinTree)RecPostOrder() {
   r := bt.root
   fmt.Print("{ ")
   bt.rPosOrd(r)
   fmt.Print("}")
}

/**************************************
 * Clear(): Função para limpar a tela *
 **************************************/
func Clear(){
   fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}

/****************************
 * main(): Função principal *
 ****************************/
func main() {
   // Declaração de variáveis
   var (
      op, x DataType
      binTree BinTree
   )
	
   for {
      Clear() // Limpa a tela
      fmt.Println("ÁRVORES BINÁRIAS DE BUSCA\n")
      fmt.Println("\t Menu Principal\n")
      fmt.Println("[ 0] Sair")
      fmt.Println("[ 1] Criar árvore")
      fmt.Println("[ 2] Imprimir árvore")
      fmt.Println("[ 3] Inserir elemento")
      fmt.Println("[ 4] Remover elemento")
      fmt.Println("[ 5] Procurar elemento")
      fmt.Println("[ 6] Percurso em largura iterativo")
      fmt.Println("[ 7] Percurso em profundide pré-ordem iterativo")
      fmt.Println("[ 8] Percurso em profundide em-ordem iterativo")
      fmt.Println("[ 9] Percurso em profundide Pós-ordem iterativo")
      fmt.Println("[10] Percurso em profundide Pré-ordem recursivo")
      fmt.Println("[11] Percurso em profundide em-ordem recursivo")
      fmt.Println("[12] Percurso em profundide Pós-ordem recursivo")
      fmt.Println("\nQual a sua opção? >> ")
      fmt.Scan(&op)
      switch op {
      case 0: {
         Clear()
         fmt.Println("Programa Encerrado!\nTecle [ENTER]")
         fmt.Scanln(&op)
         return
	 }
	 case 2: {
	    Clear()
	    fmt.Println("Árvore armazenada:\n")
	    binTree.Print(binTree.root, 0)
	    fmt.Println("\nNúmero de elementos na árvore: ", binTree.GetNumElements(), "\n")
	    fmt.Println("\nTecle [ENTER]")
	    fmt.Scanln()
	 }
	 case 3: {
	    Clear()
	    fmt.Println("Inserindo elemento\n")
	    fmt.Print("Número a inserir >> ")
	    fmt.Scanln(&x)
	    binTree.Insert(x)
	    fmt.Println("Elemento inserido. Tecle [ENTER]")
	    fmt.Scanln()
	 }
	 case 4: {
	    Clear()
	    fmt.Println("Removendo elemento\n")
	    fmt.Print("Número a remover >> ")
	    fmt.Scanln(&x)
	    binTree.Delete(x)
	    fmt.Println("Elemento Removido. Tecle [ENTER]")
	    fmt.Scanln()
	 }
	 case 5: {
	    Clear()
	    fmt.Println("Procurar elemento\n")
	    fmt.Print("Número a procurar >> ")
	    fmt.Scanln(&x)
	    a := binTree.Find(x)
	    if a == nil {
	       fmt.Println("Elemento não encontrado. Tecle [ENTER]")
	       fmt.Scanln()
	    } else {
	       fmt.Println(x, "está na árvore:", a)
	       fmt.Scanln()
	    } 
	 }
	 case 6: {
	    Clear()
	    fmt.Println("Percurso em largura iterativo\n")
	    binTree.BreadthFirst()
	    fmt.Println("\nTecle [ENTER]")
	    fmt.Scanln()
	 }
	 case 7: {
	    Clear()
	    fmt.Println("Percurso em profundidade Pré-Ordem iterativo\n")
	    binTree.ItPreOrder()
	    fmt.Println("\nTecle [ENTER]")
	    fmt.Scanln()
	 }
	 case 8: {
	    Clear()
	    fmt.Println("Percurso em profundidade Em-Ordem iterativo\n")
	    binTree.ItInOrder()
	    fmt.Println("\nTecle [ENTER]")
	    fmt.Scanln()
	 }
	 case 9: {
	    Clear()
	    fmt.Println("Percurso em profundidade Pós-Ordem iterativo\n")
	    binTree.ItPostOrder()
	    fmt.Println("\nTecle [ENTER]")
	    fmt.Scanln()
	 }
	 case 10: {
	    Clear()
	    fmt.Println("Percurso em profundidade Pré-Ordem recursivo\n")
	    binTree.RecPreOrder()
	    fmt.Println("\nTecle [ENTER]")
	    fmt.Scanln()
	 }
	 case 11: {
	    Clear()
	    fmt.Println("Percurso em profundidade Em-Ordem recursivo\n")
	    binTree.RecInOrder()
	    fmt.Println("\nTecle [ENTER]")
	    fmt.Scanln()
	 }
	 case 12: {
	    Clear()
	    fmt.Println("Percurso em profundidade Pós-Ordem recursivo\n")
	    binTree.RecPostOrder()
	    fmt.Println("\nTecle [ENTER]")
	    fmt.Scanln()
	 }
	 default: {
	    fmt.Println("Opção Inválida!\nTecle [ENTER]")
	    fmt.Scanln(&op)
	 }
      }
   }
}
