package main

import "fmt"

/*
- Testes devem:
    - ficar num arquivo cuja terminação seja _test.go
    - ficar na mesma package que o código a ser testado
    - ficar em funções com nome "func TestNome(*testing.T)"
- Para rodar os testes:
    - go test
    - go test -v
- Para falhas, utilizamos t.Error(), onde a maneira idiomática é algo do tipo "expected: x. got: y."

- Outra maneira é fazer testes como exemplos.
- Estes exemplos são os mesmos que aparecem na documentação.
- Para exemplos o formato é "func ExampleFuncao()"
- Deve haver um comentário "// Output: resultado", que é o que será testado
- Para visualizar seu exemplo na documentação, fazemos o de sempre:
    - godoc -http :8080
- Tanto para testes quanto para exemplos podemos utilizar: go test ./...

*/

func main() {
	x := Soma(1, 2, 3)
	y := Multiplica(10, 10)
	fmt.Println(x, y)
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
		// total += v
	}
	return total
}

func Soma2(i ...int) int {
	total := 0
	for _, v := range i {
		//total = total + v
		total += v
	}
	return total
}

func Multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func Multiplica2(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}