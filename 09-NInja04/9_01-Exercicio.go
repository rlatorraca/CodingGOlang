package main

import "fmt"

/*
- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
 */


func main(){

	array := [5]int{10, 20, 30, 40, 50}

	for i, v := range array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", array)

	array2 := [5]int{}
	array2[0] = 10
	array2[1] = 11
	array2[2] = 21
	array2[3] = 31
	array2[4] = 41


	for i, v := range array2 {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", array)
}