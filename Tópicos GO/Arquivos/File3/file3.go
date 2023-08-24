package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// abrindo arquivo de entrada no buffer infile
	infile, ferror := os.Open("arquivoin.txt")
	// verifica possivel erro
	if ferror != nil { // ferror != NULL
		fmt.Println("ERRO[1]: erro ao abrir arquivo in")
		panic(ferror) // termina execucao imediatamento, erros fatais
	}

	// defer adia a execucao de uma funcao ate o termino da funcao que a contem
	// funcao anonima para fechar o arquivo
	defer func() {
		if err := infile.Close(); err != nil {
			fmt.Println("ERRO[2]: erro ao fechar arquivo in.")
			panic(err)
		}
	}()

	// criando arquivo de saida
	outfile, ferror := os.Create("arquivoout.txt")
	if ferror != nil {
		fmt.Println("ERRO[3]: erro ao criar arquivo out.")
		panic(ferror)
	}

	defer func() {
		if err := outfile.Close(); err != nil {
			fmt.Println("ERRO[4]: erro ao fechar arquivo out.")
			panic(err)
		}
	}()

	// criando um buffer para conter o que sera lido do arquivo in
	buf := make([]byte, 1024)
	for {
		// leitura do arquivo in
		reader, ferror := infile.Read(buf)
		if ferror != nil && ferror != io.EOF {
			fmt.Println("ERRO[5]: erro na leitura do arquivo in.")
			panic(ferror)
		}
		// fim do arquivo
		if reader == 0 {
			break
		}

		// escrevendo o buffer no arquivo our
		if _, ferror := outfile.Write(buf[:reader]); ferror != nil {
			fmt.Println("ERRO[6]: erro ao escrever no arquivo out.")
			panic(ferror)
		}
	}

}
