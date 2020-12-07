package main

import "fmt"

/*
- Pode-se usar uma função como retorno de uma função
- Declaração: func f() return
- Exemplo: func f() func() int { [...]; return func() int{ return [int] } }
    - ????: fmt.Println(f()())
*/

func main() {

	x := retornaumafuncao()

	y := x(3)

	fmt.Println(y)

	fmt.Println(retornaumafuncao()(3))

}

// Funcao que retorna um FUNCAO que essa funcao retorna um int
func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
