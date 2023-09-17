/*
- Usando uma literal composta:
  - Crie uma slice de tipo int
  - Atribua 10 valores a ela

- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
*/
package main

import (
	"fmt"
)

func Exercicio22() {
	fmt.Println(`Usando uma literal composta:`)
	fmt.Println("")
	slice := []int{}
	fmt.Printf("- Crie uma slice de tipo int:")
	fmt.Println("")
	fmt.Printf("%T", slice)
	fmt.Println("")
	fmt.Println("")

	fmt.Printf("- Atribua 10 valores a ela: \n")
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)
	fmt.Println("")

	fmt.Printf("- Utilize range para demonstrar todos estes valores.\n")
	for v := range slice {
		fmt.Println(slice[v])
	}
	fmt.Println("")
	fmt.Printf("- E utilize format printing para demonstrar seu tipo.\n")
	fmt.Printf("Type of slice: %T", slice)
}
