/*
Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/
package main

import (
	"fmt"
	// os
	"strings"
)

/*
	JOGO: Acerte a fruta!
*/

func Exercicio19() {

	// Declara e atribui uma string a variável
	fruta := "laranja"

	// Converte a string para caixa baixa (minúsculo)
	fruta = strings.ToLower(fruta)

	switch {
	case fruta == "laranja":
		fmt.Println("Essa é a fruta!")
	default:
		fmt.Println("Essa não é a fruta!")
	}
}
