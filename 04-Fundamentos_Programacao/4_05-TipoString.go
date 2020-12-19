package main

/*
- Strings são sequencias de bytes.
- Imutáveis.
- Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
- Na prática:
    - %v (valor) %T (tipo) %b(Binario)
    - Raw string literals
    - Conversão para slice of bytes: []byte(x)
    - %#U (unicode) , %#x (Hexadecimal)
    - Go Playground: https://play.golang.org/p/dt2x1ies5b & https://play.golang.org/p/PpDnspiyA_7
- https://blog.golang.org/strings
 */
import (
	"fmt"
)

func main() {
	s := "ascii éøâ 香"
	sb :=[]byte(s)
	t := "Hi, Brazil"

	fmt.Printf("%v %T\n", sb, sb)
	fmt.Printf("%v %T\n", t, t)

	// Devolve CARACTER por CARACTER (usando RANGE)
	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println()

	// Deveolve BYTE a BYTE (For tradicional))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}

}
