/*
- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.
*/
package main

import (
	"fmt"
)

func Exercicio8() {

	const constanteTipada int = 10
	const constanteNaoTipada = 10

	fmt.Println("Constante tipada:", constanteTipada)
	fmt.Println("Constante não tipada:", constanteNaoTipada)
}
