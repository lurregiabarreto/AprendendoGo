/*
- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
*/
package main

import "fmt"

func Exercicio16() {

	// Loop que começa em 10 e termina em 100
	for i := 10; i <= 100; i++ {

		// Imprime o resto da divisão por 4
		fmt.Println(i % 4)
	}
}
