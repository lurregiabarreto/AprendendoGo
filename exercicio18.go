/*
- Utilizando a solução anterior, adicione as opções else if e else.
*/
package main

import (
	"fmt"
	//"os"
	"strings"
)

/*
	JOGO: Acerte a fruta!
*/

func Exercicio18() {

	fruta := "cerveja"
	fruta = strings.ToLower(fruta)

	// Condição que verifica qual é a fruta da variável.
	if fruta == "maçã" {
		fmt.Println("Errado, maçã não é a fruta! =)")
	} else if fruta == "amora" {
		fmt.Println("Errado, amora não é a fruta! :D")
	} else if fruta == "cereja" {
		fmt.Println("Errado, cereja não é a fruta! :O")
	} else if fruta == "laranja" {
		fmt.Println("Acertou! A fruta é laranja! :)")
	} else {
		fmt.Println("Errado,", fruta, "não é a fruta!")
	}
}
