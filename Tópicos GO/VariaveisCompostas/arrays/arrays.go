package main

import "fmt"

func main() {
	var (
		idades [5]int
		nomes  = [5]string{"Luiz", "Eduarda", "Caio", "Heloisa", "Marcos"}
		dados  [5]string
		i      int
	)

	for i = 0; i < len(idades); i++ {
		fmt.Print("Informe uma idade para ", nomes[i], ": ")
		fmt.Scan(&idades[i])
		dados[i] = fmt.Sprint("Nome: ", nomes[i], " Idade: ", idades[i])
	}

	i = 0
	for i < len(dados) {
		fmt.Println(dados[i])
		i++
	}

}
