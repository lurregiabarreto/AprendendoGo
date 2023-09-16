/*
- Usando uma literal composta:
  - Crie um array que suporte 5 valores to tipo int
  - Atribua valores aos seus índices

- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
*/
package main

import (
	"fmt"
)

var array = [5]int{}

func Exercicio21() {
	fmt.Println(`Usando uma literal composta:`)
	fmt.Println("")

	fmt.Println(`- Crie um array que suporte 5 valores to tipo int: `)
	fmt.Println(array)
	array[0] = 1
	array[1] = 2
	array[2] = 1337
	array[3] = 345
	array[4] = 45353

	fmt.Println("")
	fmt.Println(`- Atribua valores aos seus índices`)

	for i, v := range array {
		fmt.Printf("Indice: %v => Valor: %v\n", i, v)
	}
	fmt.Println("")
	fmt.Println(`- Utilize range e demonstre os valores do array.`)

	for i := range array {
		fmt.Println(array[i])
	}

	fmt.Println("")
	fmt.Println(`- Utilizando format printing, demonstre o tipo do array.`)

	fmt.Printf("Type array: %T", array)
}
