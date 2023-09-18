/*
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
  - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
  - Do quinto ao último item do slice (incluindo o último item!)
  - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
  - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
  - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
*/
package main

import (
	"fmt"
)

func Exercicio23() {
	slice := []int{}
	fmt.Printf("- Atribua 10 valores a ela: \n")
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)
	fmt.Println("- Do primeiro ao terceiro item do slice (incluindo o terceiro item!)")
	fmt.Println(slice[:3])
	fmt.Println("- Do quinto ao último item do slice (incluindo o último item!)")
	fmt.Println(slice[4:])
	fmt.Println("- Do segundo ao sétimo item do slice (incluindo o sétimo item!)")
	fmt.Println(slice[1:7])
	fmt.Println("- Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)")
	fmt.Println(slice[2:9])
	fmt.Println("- Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item")
	fmt.Println(slice[2:len(slice[1:])])

}
