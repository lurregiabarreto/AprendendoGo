/*
- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
*/
package main

import (
	"fmt"
)

func Exercicio6() {
	x := 13878493
	fmt.Printf("O valor de %v em binário é: %b \n", x, x)
	fmt.Printf("O valor de %v em decimal é: %d \n", x, x)
	fmt.Printf("O valor de %v em octagonal é: %o \n", x, x)
	fmt.Printf("O valor de %v em hexadecimal é: %x \n", x, x)

}
