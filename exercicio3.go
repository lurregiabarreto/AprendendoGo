/*  1.0 Em package-level scope, atribua os seguintes valores às variáveis:
1. para "x" atribua 42
2. para "y" atribua "James Bond"
3. para "z" atribua true

2.0 Na função main:
1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
2. Demonstre a variável "s".
*/

package main

import (
	"fmt"
)

var a int
var b string
var c bool

func Exercicio3() {
	a = 42
	b = "James Bond"
	c = true
	s := fmt.Sprintf("O valor de A, B e C nas respectiva ordem são: %v, %v e %v", a, b, c)

	fmt.Println("Na váriavel S contem: ", s)
}
