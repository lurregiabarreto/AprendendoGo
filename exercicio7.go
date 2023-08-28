/*
- Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
- Demonstre estes valores.
*/

package main

import (
	"fmt"
)

func Exercicio7() {

	x := 10
	y := 10

	igualdade := x == y
	diferenca := x != y
	menorIgual := x <= y
	menor := x < y
	maiorIgual := x >= y
	maior := x > y

	fmt.Println("O valor da == entre x e y = ", igualdade)
	fmt.Println("O valor da != entre x e y = ", diferenca)
	fmt.Println("O valor da <= entre x e y = ", menorIgual)
	fmt.Println("O valor da < entre x e y = ", menor)
	fmt.Println("O valor da >= entre x e y = ", maiorIgual)
	fmt.Println("O valor da >= entre x e y = ", maior)
}
