/**********************************************************************
 * HashTable.go: Implementa uma tabela hash simples utilizando arrays *
 **********************************************************************/
package main
 
import(
	"fmt"
)

/*********************************
 * Declaração do tipo HashTable  *
 *********************************/
type HashTable struct {
	tableSize int8
	bucketSize int8
	table[][] int8
}

/***********************************************
 * Função de inicialização da classe HashTable *
 ***********************************************/
func (ht *HashTable) Init(tsize int8, bsize int8) {
	// seta os atributos da classe
	ht.tableSize = tsize
	ht.bucketSize = bsize
	
	// Aloca a tabela hash
   ht.table = make([][]int8, ht.tableSize)
   for i := range ht.table {
        ht.table[i] = make([]int8, ht.bucketSize)    
   }
   
   // Preenche com os valores iniciais
   for i := range ht.table {
		for j := range ht.table[i] {
         ht.table[i][j] = -1
		}
	}
}

/****************************************
 * Função Print para a classe HashTable *
 ****************************************/
func (ht *HashTable) Print() {
	for i := range ht.table {
			fmt.Print(ht.table[i])
			fmt.Println()
	}
}

/*****************************************************************
 * Função hash para calcular o índice h para a chave k na tabela *
 *****************************************************************/
 func (ht *HashTable) h(k int8) int8 {
	// calcula o índice h na tabela correspondente a chave k
   if k < 0 {
		fmt.Print("Erro 1: Chave k é negativa.\nRetornando valor -1")
      return(-1)
   } else {
      return(k%ht.tableSize)
   }
 }
 
/**********************
 * Função de inserção *
 **********************/
 func (ht *HashTable) Insert(k int8){
   i := ht.h(k)
   if i < 0 {
      panic("Erro: índice inválido")
   } else {
		var j int8
      for (int8(j) < ht.bucketSize)&&(ht.table[i][j] != -1){
			j++
		}
      if (int8(j) < ht.bucketSize) && (ht.table[i][j] == -1) {
         ht.table[i][j] = k
      } else {
         panic("Erro 2: Bucket cheio!")
      }
	}
}

/*********************
 * Função de remoção *
 *********************/
func (ht *HashTable) Remove(k int8){
	var (
		i int8 = ht.h(k)
		j int8 
	)
   if i < 0 {
		panic("Erro: Índice inválido!")
   } else {  
		achou := false
      for ;(!achou) && (j < ht.bucketSize); j++ {
			if k == ht.table[i][j] {
				achou = true
				ht.table[i][j] = -1
         } else {
				j++
			}
      }
   }
}

/*******************
 * Função de busca *
 *******************/
func (ht *HashTable) search(k int8) (pos int8) {
	var (
		i int8 = ht.h(k)
		j int8
		achou bool
	)
	
   if i < 0 {
		pos = -1
	} else {
		for ;(!achou) && (j < ht.bucketSize); j++{
			if k == ht.table[i][j] {
				achou = true
				pos = k
			}
		}
	}
	return pos
}

/********************
 * Função Principal *
 ********************/
func main (){
	// Declaraç˜ão de Variáveis
	var (
		ht HashTable
	)
	
	fmt.Print("Armazenamento em HashTable\n\n")
	fmt.Print("Inicialização da tabela hash:")
	ht.Init(10, 2)
	ht.Print()
	fmt.Print("Inserindo 1:\n")
	ht.Insert(1)
	fmt.Print("Inserindo 2:\n")
	ht.Insert(2)
	ht.Print()
	fmt.Print("Inserindo 4:\n")
	ht.Insert(4)
	fmt.Print("Inserindo 3:\n")
	ht.Insert(3)
	fmt.Print("Inserindo 8:\n")
	ht.Insert(8)
	fmt.Print("Inserindo 5:\n")
	ht.Insert(5)
	fmt.Print("Inserindo 6:\n")
	ht.Insert(6)
	fmt.Print("Inserindo 7:\n")
	ht.Insert(7)
	fmt.Print("Inserindo 16:\n")
	ht.Insert(16)
	fmt.Print("Inserindo 9:\n")
	ht.Insert(9)
	fmt.Print("Inserindo 10:\n")
	ht.Insert(10)
	fmt.Print("Inserindo 11:\n")
	ht.Insert(11)
	fmt.Print("Removendo 2:\n")
	ht.Remove(2)
	ht.Print()
} 
