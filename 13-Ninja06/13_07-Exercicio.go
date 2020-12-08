package main

import "fmt"

/*
- Atribua uma função a uma variável.
- Chame essa função.
 */



func main() {
	x:="Jorginho"

	y:= func(x string){
		fmt.Println("O nome de X :" , x)
	}
	y(x)
}