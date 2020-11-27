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

type cachorroQuente int

var x5 cachorroQuente

var y5 int

func main() {
	fmt.Println(x5)
	fmt.Printf("%T\n", x5)
	x5 = 42
	fmt.Println(x5)
	fmt.Println("↑ foi x.\n↓ é y!")
	y5 = int(x5)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)

}
