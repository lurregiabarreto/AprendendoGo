/*
  - Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
  - Por exemplo:
    65
    U+0041 'A'
    U+0041 'A'
    U+0041 'A'
    66
    U+0042 'B'
    U+0042 'B'
    U+0042 'B'
    ...e por aí vai.
*/
package main

import "fmt"

func Exercicio13() {

	// Loop que começa em 60 e termina em 90
	for i := 65; i <= 90; i++ {
		// Imprime o número do UNICODE
		fmt.Println(i)

		// Loop que começa em 1 e termina no 3
		for x := 1; x <= 3; x++ {
			// Imprime o UNICODE
			fmt.Printf("%#U\n", i)
		}

		// Quebra de linha
		fmt.Print("\n")
	}
}
