/*
- Crie um map com key tipo string e value tipo []string.
  - Key deve conter nomes no formato sobrenome_nome
  - Value deve conter os hobbies favoritos da pessoa

- Demonstre todos esses valores e seus índices.
*/
package main

import (
	"fmt"
)

func Exercicio28() {

	y := map[string][]string{
		"ana":     []string{"hacking", "31447", "snow"},
		"matheus": []string{"malandro", "theus", "math"},
		"andré":   []string{"viaja", "programa", "qualquer"},
	}
	for i, v := range y {
		fmt.Println(i)
		for z, b := range v {
			fmt.Println("\t", z, "\t-", b)
		}
	}

}
