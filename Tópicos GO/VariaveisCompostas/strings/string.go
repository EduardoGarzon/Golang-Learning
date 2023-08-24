package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	var (
		A, B string
	)

	fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)

	fmt.Println("\033[36mTrabalhando com Strings\033[0m") // texto ciano

	// LENDO STRINGS
	fmt.Println("1. Leitura de Strings:")
	fmt.Print("\tDigite uma frase: ")
	scanner := bufio.NewReader(os.Stdin) // le a entrada do console
	A, _ = scanner.ReadString('\n')      // le ate o \n e descarta o segundo retorno para erros _
	A = strings.Trim(A, "\n")            // remover \n da string

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// ESCREVENDO STRINGS
	fmt.Println("2. Escrita de strings:")
	fmt.Println("\tString lida:", A)

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// ALTERANDO STRINGS
	fmt.Println("3. Alteração de strings:")
	//S[0] = 'A'  // Isto não funciona!
	//S[0] = "A" // Nem isso
	A = "A" + A // Isso Funciona! (concatenação)
	fmt.Println("\tString Alterada  =", A)

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// COMPRIMENTO DE STRINGS
	fmt.Println("4. Comprimento de strings:")
	r := "Luiz"
	fmt.Println("\tlen(A) =", len(r))                        // bytes
	fmt.Println("\tcaracteres =", utf8.RuneCountInString(r)) // caracteres unicode

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// COMPARANDO STRINGS
	fmt.Println("5. Comparação de Strings:")
	fmt.Print("\tdigite uma frase: ")
	scanner = bufio.NewReader(os.Stdin)
	B, _ = scanner.ReadString('\n')
	B = strings.Trim(B, "\n")
	fmt.Println("\tString lida = ", B)
	fmt.Println("\tString A = ", A)
	fmt.Println("\tString B = ", B)
	if strings.Compare(A, B) == 0 {
		fmt.Println("\tA e B são iguais.")
	} else {
		fmt.Println("\tA e B são diferentes.")
	}

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// PROCURA POR SUBSTRINGS
	fmt.Println("6. Procura por substrings:")
	if strings.Contains(B, A) {
		fmt.Println("\t", A, "está em", B)
	} else {
		fmt.Println("\t", A, "não está em", B)
	}
	if strings.Contains(A, B) {
		fmt.Println("\t", B, "está em", A)
	} else {
		fmt.Println("\t", B, "não está em", A)
	}

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// PROCURA POR CONJUNTO DE CARACTERES EM UMA STRING
	fmt.Println("7. Procura por conjuntos de caracteres:")
	if strings.ContainsAny(A, "cef") {
		fmt.Println("\t", A, "contém algum dos caracteres", "cef")
	} else {
		fmt.Println("\t", A, "não contém algum dos caracteres", "cef")
	}

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// PROCURA POR CARACTERES (RUNAS) ESPECIFICAS
	fmt.Println("8. Procura por caracteres (runas):")
	if strings.ContainsRune(A, rune('a')) {
		fmt.Println("\t", A, "contém a runa ", "a")
	} else {
		fmt.Println("\t", A, "não contém a runa", "a")
	}

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// Procura e conta substrings
	fmt.Println("9. Procura e conta substrings:")
	i := strings.Count(A, B)
	fmt.Println("\t", B, "aparece", i, "vezes em", A)

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// Recortar strings
	fmt.Println("10. Recortar strings")
	b, a, f := strings.Cut(A, " ")
	fmt.Printf("\tb: %q\n\ta: %q\n\tf: %t \n\n", b, a, f)

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// Indexação de strings
	fmt.Println("11. indexação de strings")
	fmt.Println("\tstrings.Index(A, B) = ", strings.Index(A, B))
	fmt.Println("\tstrings.IndexAny(A, \"aeiou\") = ", strings.IndexAny(A, "aeiou"))
	fmt.Println("\tstrings.IndexByte(A, 'b') = ", strings.IndexByte(A, 'b'))
	fmt.Println("\tstrings.IndexRune(A, 'c') = ", strings.IndexRune(A, 'c'))

	fmt.Println("\tstrings.LastIndex(A, B) = ", strings.LastIndex(A, B))
	fmt.Println("\tstrings.LastIndexAny(A, \"aeiou\") = ", strings.LastIndexAny(A, "aeiou"))
	fmt.Println("\tstrings.LastIndexByte(A, 'b') = ", strings.LastIndexByte(A, 'b'))
	fmt.Println("\tstrings.LastIndexRune(A, 'c') = ", strings.IndexRune(A, 'c'))

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// Indexação de strings com strings.IndexFunc e strings.LastIndexFunc
	fmt.Println("12. indexação de strings com strings.IndexFunc")
	fmt.Println("\tstrings.IndexFunc(A, unicode.IsNumber)", strings.IndexFunc(A, unicode.IsNumber))
	fmt.Println("\tstrings.IndexFunc(A, unicode.IsPunct)", strings.IndexFunc(A, unicode.IsPunct))
	fmt.Println("\tstrings.IndexFunc(A, unicode.IsLetter)", strings.IndexFunc(A, unicode.IsLetter))
	fmt.Println("\tstrings.IndexFunc(A, unicode.IsSpace)", strings.IndexFunc(A, unicode.IsSpace))
	fmt.Println("\tstrings.IndexFunc(A, unicode.IsUpper)", strings.IndexFunc(A, unicode.IsUpper))

	fmt.Println("\tstrings.LastIndexFunc(A, unicode.IsNumber)", strings.LastIndexFunc(A, unicode.IsNumber))
	fmt.Println("\tstrings.LastIndexFunc(A, unicode.IsPunct)", strings.LastIndexFunc(A, unicode.IsPunct))
	fmt.Println("\tstrings.LastIndexFunc(A, unicode.IsLetter)", strings.LastIndexFunc(A, unicode.IsLetter))
	fmt.Println("\tstrings.LastIndexFunc(A, unicode.IsSpace)", strings.LastIndexFunc(A, unicode.IsSpace))
	fmt.Println("\tstrings.LastIndexFunc(A, unicode.IsUpper)", strings.LastIndexFunc(A, unicode.IsUpper))

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	// Funções strings.Join e strings.Fields
	fmt.Println("13. Exemplo de strings.Join e strings.Fields")
	split := strings.Fields(A) // cria um slice (lista) com as substrings do parâmetro A
	fmt.Println("\t", split)
	join := strings.Join(split, "_") // cria uma nova string com os elementos da slice e o separador '_' no lugar dos espaços
	fmt.Println("\t", join)

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	fmt.Println("14. Exemplo de substituição usando strings.Replace e strings.ReplaceAll")
	// Substituição em strings com strings.Replace e strings.ReplaceAll
	fmt.Println(strings.Replace(A, "u", "~", 1))
	fmt.Println(strings.Replace(A, "u", "~", 2))
	fmt.Println(strings.Replace(A, B, "@@", -1))
	fmt.Println(strings.ReplaceAll(A, B, "###"))

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

	fmt.Println("15. Como varrer uma string")
	// Como varrer uma string
	r2 := []rune(A)
	S3 := string(r2)
	// lr := len(r)
	for i := range A {
		fmt.Println("S[", i, "] = ", string(rune(A[i])))
	}

	fmt.Println("S3 = ", strings.ToUpper(S3))

	fmt.Print("Tecle [ENTER]")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")
}
