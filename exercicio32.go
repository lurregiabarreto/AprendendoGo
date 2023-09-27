/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/
package main

import (
	"fmt"
)

type people struct {
	nome, sobrenome string
	sabores         []string
}

func Exercicio32() {

	m := make(map[string]people)

	m["barbosa"] = people{
		nome:      "eduardo",
		sobrenome: "barbosa",
		sabores:   []string{"abacate", "chocolate", "morango"},
	}

	m["Rocha"] = people{
		nome:      "Maria",
		sobrenome: "Rocha",
		sabores:   []string{"Maracuja", "Morango"},
	}

	for _, v := range m {
		fmt.Println(v.nome)
		for _, sabor := range v.sabores {
			fmt.Println("\t", sabor)

		}

	}

}
