/*
- Comece com essa slice:
  - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

- Utilizando slicing e append, crie uma slice y que contenha os valores:
  - [42, 43, 44, 48, 49, 50, 51]
*/
package main

import (
	"fmt"
)

func Exercicio25() {
	fmt.Printf(`- Comece com essa slice:
    	- x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}`)
	fmt.Println("")
	fmt.Println("")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println("Utilizando slicing e append, crie uma slice y que contenha os valores:")
	fmt.Println("[42, 43, 44, 48, 49, 50, 51]")
	fmt.Println("")
	z := x[:3]
	b := x[6:]
	y := append(z, b...)
	fmt.Println("FINAL:", y)

}
