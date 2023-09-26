/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
  - Nome
  - Sobrenome
  - Sabores favoritos de sorvete

- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/
package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func Exercicio31() {
	pessoa1 := pessoa{
		nome:      "Alberto",
		sobrenome: "Rocha",
		sabores:   []string{"Baunilha", "Chocolate"},
	}

	pessoa2 := pessoa{
		nome:      "Maria",
		sobrenome: "Rocha",
		sabores:   []string{"Maracuja", "Morango"},
	}

	fmt.Println(pessoa1.nome)
	imprimirSabores(pessoa1.sabores)
	fmt.Println(pessoa2.nome)
	imprimirSabores(pessoa2.sabores)

}

func imprimirSabores(condicao []string) {
	for _, y := range condicao {
		fmt.Println("\t", y)
	}
}
