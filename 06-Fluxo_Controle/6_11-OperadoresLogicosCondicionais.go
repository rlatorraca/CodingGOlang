package main

/*
- && (E)
- || (OU)
- !   (negacao)
- Go Playground: https://play.golang.org/p/MFwrt93xlc
- Qual o resultado de fmt.Println...
    - true && true
    - true && false
    - true || true
    - true || false
    - !true
 */

import (
	"fmt"
)

func main() {

	x := 6

	if !(x%2 == 0) && x%3 == 0 {
		fmt.Println("é múltiplo de dois e tambem de treis")
	}

	if x%2 == 0 || x%3 == 0 {
		fmt.Println("é múltiplo de dois ou de treis")
	}

	if true && true {
		fmt.Println("true && true")
	}
	if true && false {
		fmt.Println("true && false")
	}

	if false && true {
		fmt.Println("false && true")
	}

	if false && false {
		fmt.Println("false && false")
	}

	if true || true {
		fmt.Println("true || true")
	}

	if true || false {
		fmt.Println("true || false")
	}

	if false || true {
		fmt.Println("false || true")
	}

	if false || false {
		fmt.Println("false || false")
	}

	if !true {
		fmt.Println("!true")
	}

}
