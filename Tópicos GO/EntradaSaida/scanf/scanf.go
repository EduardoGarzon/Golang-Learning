/*
	funciona igual C e C++
*/

package main

import (
	"fmt"
	//"os"
	//"bufio"
)

func main() {
	var nome string
	var idade int

	fmt.Scanf("%s%*c", &nome)
	fmt.Printf("%s\n", nome)

	//reader := bufio.NewReader(os.Stdin)
	//reader.ReadString('\n')

	fmt.Scanf("%d", &idade)
	fmt.Println(idade)
}
