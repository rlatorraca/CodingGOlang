package main

import "fmt"

/*
- Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
    - [42, 43, 44, 48, 49, 50, 51]
 */


func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

		x = append(x[:3], x[len(x)-4:]...)
	//x = append(x[:3], x[6:]...)

	//[42, 43, 44, 48, 49, 50, 51]

	fmt.Println(x)
}