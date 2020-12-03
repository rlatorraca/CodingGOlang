package main

import "fmt"

/*
- Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
 */


func main(){

	slice := []int{10, 20, 30, 40, 50, 1, 2, 3, 4, 5}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slice)

}