/*
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/
package main

import "fmt"

func Exercicio20() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "volei":
		fmt.Print("amo jogar volei!")

	case "basquete":
		fmt.Print("amo jogar basquete!")

	case "futebol":
		fmt.Print("amo jogar futebol!")
	}
}
