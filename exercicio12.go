/*
Laços de repetição
Põe na tela: todos os números de 1 a 10000.
*/
package main

import "fmt"

// todos numero de 1 a 10000

func Exercicio12() {
	n := 1
	for {
		if n > 10000 {
			break
		}
		fmt.Println("Numero:", n)
		n++
	}
}
