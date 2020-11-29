package main

/*
- Operação módulo: %
- For: break
- For: continue
 */


import (
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		if i % 2 != 0 {
			// faz isso quando o número não é par (IMPAR)
			continue // Vai para proxima iteracao
		}
		fmt.Println(i)
	}
	fmt.Println()
	for i := 0; i < 20; i++ {
		if i == 3 {
			// faz isso quando o número não é par
			continue // Vai para proxima iteracao (quebra apenas a ITERACAO e nao o LOOP )
		}
		fmt.Println(i)
	}
}