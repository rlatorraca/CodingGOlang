package main

/*
- Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.

- https://splice.com/blog/iota-elegant-...
https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

 */

import (
	"fmt"
)

const (
	/*
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
	*/
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB
	GB
	TB
)

func main() {
	x := 234
	y := x << 1
	z := y << 1
	w := z >> 2

	fmt.Printf("Valor binário :%b , %d\n", x, x)
	fmt.Printf("Valor binário :%b , %d\n", y, y)
	fmt.Printf("Valor binário :%b , %d\n", z, z)
	fmt.Printf("Valor binário :%b , %d\n", w, w)

	fmt.Println()

	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
