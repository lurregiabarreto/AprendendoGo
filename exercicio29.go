/*
- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
*/
package main

import (
	"fmt"
)

func Exercicio29() {
	mepezin := map[string][]string{"cores": []string{"verde", "azul", "preto"}}

	mepezin["tools"] = []string{"gobuster", "dirbuster", "set"}

	fmt.Println(mepezin)

	for i, v := range mepezin {
		fmt.Println(i, v)
	}
}
