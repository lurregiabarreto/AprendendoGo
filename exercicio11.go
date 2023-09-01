/*
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
*/
package main

import (
	"fmt"
)

const (
	_ = 2021 + iota
	primeiroAno
	segundoAno
	terceiroAno
	quartoAno
)

func Exercicio11() {

	// Imprime as constantes
	fmt.Println(primeiroAno, segundoAno, terceiroAno, quartoAno)
}
