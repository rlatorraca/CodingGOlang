package main

/*
- Agora vamos explorar os tipos de maneira mais detalhada. golang.org/ref/spec. A começar pelo bool.
- O tipo bool é um tipo binário, que só pode conter um dos dois valores: true e false. (Verdadeiro ou falso, sim ou não, zero ou um, etc.)
- Sempre que você ver operadores relacionais, o resultado da expressão será um valor booleano.
- Booleans são fundamentais nas tomadas de decisões em lógica condicional, declarações switch, declarações if, fluxo de controle, etc.
- Na prática:
    - Zero value
    - Atribuindo um valor
    - Bool como resultado de operadores relacionais
- Go Playground: https://play.golang.org/p/7joj615nZw

*/
import "fmt"

var x bool

func main() {
	fmt.Println(x) // zero value
	x = true
	fmt.Println(x) // valor atribuido
	x = 10 == 100
	y := 10 == 100
	z := 10 > 100
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

