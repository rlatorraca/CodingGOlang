package main
/*
- If, else.
- If, else if, else.
- If, else if, else if, ..., else.
 */

import (
	"fmt"
)

func main() {
	if x := 50; x > 100 {
		fmt.Println("X é maior que cem")
	} else if x < 10 {
		fmt.Println("X é menor que 10")
	} else {
		fmt.Println("X não é menor que 10 nem maior que cem")
	}
}
