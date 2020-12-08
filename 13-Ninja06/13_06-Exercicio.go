package main

import "fmt"

/*
- Crie e utilize uma função anônima.
 */



func main() {
	x:="Joaozinho"

	func(x string){
		fmt.Println("O nome de X :" , x)
	}(x)
}