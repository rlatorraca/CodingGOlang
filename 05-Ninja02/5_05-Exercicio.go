package main

/*
- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
 */


import (
	"fmt"
)

func main() {
	x := `isto
	é
		uma coisa
			muito doida`
	fmt.Println(x)
}