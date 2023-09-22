/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
  - "Nome", "Sobrenome", "Hobby favorito"

- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
package main

import (
	"fmt"
)

func Exercicio27() {
	slice := [][][]string{{
		{"ana", "anake", "Hacking 1228"},
		{"Eduardo", "Barbosa", "Programmer"},
		{"Camila", "Ramos", "Praiera"}}}

	for _, name := range slice {
		fmt.Println("Nome:", name[0][0])
		fmt.Println("Sobrenome:", name[0][1])
		fmt.Println("Hobby favorito:", name[0][2])
		fmt.Println("-------------------------------")
		fmt.Println("Nome:", name[1][0])
		fmt.Println("Sobrenome:", name[1][1])
		fmt.Println("Hobby favorito:", name[1][2])
		fmt.Println("-------------------------------")
		fmt.Println("Nome:", name[2][0])
		fmt.Println("Sobrenome:", name[2][1])
		fmt.Println("Hobby favorito:", name[2][2])
	}
}
