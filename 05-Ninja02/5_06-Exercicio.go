package main

import "fmt"

/*
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.

*/

const (
	_ = 2018 + iota * 4 // Base
	b
	c
	d
	e
)

func main() {
	// Proximos anos de Copa do Mundo
	fmt.Println(b, c, d, e)
}
