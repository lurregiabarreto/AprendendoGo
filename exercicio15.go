/*
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/
package main

import "fmt"

func Exercicio15() {

	// Declara e atribui valores inteiros a variáveis
	anoNascimento := 1997
	anoAtual := 2023

	// Loop infinito
	for {

		// Condição para checar se o anoNascimento é menor ou igual ao anoAtual
		if anoNascimento <= anoAtual {

			// Imprime o loop
			fmt.Println("Ano:", anoNascimento)

			// Adiciona +1 a variável
			anoNascimento++

		} else {

			// Interrompe o loop caso a condição seja falso
			break
		}
	}
}
