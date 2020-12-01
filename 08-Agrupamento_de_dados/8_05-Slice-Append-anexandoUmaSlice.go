package main

/*
 Effective Go: append (package builtin)
- x = append(slice, ...values)
- x = append(slice, slice...)

- Todd: unfurl → desdobrar, desenrolar  (esses 3 pontos) "..."
- Nome oficial: enumeration
 */


import (
	"fmt"
)

func main() {

	umaSlice := []int{1, 2, 3, 4}
	outraSlice := []int{9, 10, 11, 12}

	fmt.Println(umaSlice)

	umaSlice = append(umaSlice, 5, 6, 7, 8)

	fmt.Println(umaSlice)

	umaSlice = append(umaSlice, outraSlice...) // "... " -> transforma em primitivo (int no caso) , ooutraSlice *que é um "Slice de int"

	fmt.Println(umaSlice)
}