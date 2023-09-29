/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/
package main

import (
	"fmt"
)

func Exercicio34() {

	anon := struct {
		nome  map[string]string
		valor []int
	}{
		map[string]string{"hello": "string"},
		[]int{1, 2, 3, 4, 5},
	}

	fmt.Println(anon)

}
