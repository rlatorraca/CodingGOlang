package main
/*
- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
*/
import (
	"fmt"
)

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Println("Para o valor de X ", x, " o resto da divisao por 4 é : ",  x % 4)
	}
}
