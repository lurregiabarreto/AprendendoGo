/*
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/
package main

import "fmt"

func Exercicio14() {

	// Declara e atribui valores inteiros a variáveis
	anoNascimento := 1997
	anoAtual := 2023

	// Loop que começa em 1997 e termina em 2023
	for anoNascimento <= anoAtual {

		// Imprime o loop
		fmt.Println("Ano:", anoNascimento)

		// Adiciona +1 a variável
		anoNascimento++
	}
}
