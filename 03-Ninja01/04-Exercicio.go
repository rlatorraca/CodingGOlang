package main

/*
- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"
- Para os aventureiros: https://golang.org/ref/spec#Types
- Agora já somos quase ninjas nível 1!
- Solução: https://play.golang.org/p/snm4WuuYmG
*/
import (
	"fmt"
)

type hotdog int

var x4 hotdog

func main() {
	fmt.Println(x4)
	fmt.Printf("%T\n", x4)
	x4 = 42
	fmt.Println(x4)
}
