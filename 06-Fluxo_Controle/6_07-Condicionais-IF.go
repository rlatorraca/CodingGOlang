package main
/*
- If: bool
- If: o operador não → "!"
- If: declaração de inicialização
 */


import (
	"fmt"
)

func main() {
	if x := 10; !(x > 100) {
		fmt.Println("Hello, playground")
	}
	fmt.Println()
	//OR
	y := 10;
	if  y > 100 {
		fmt.Println("Hello, playground")
	} else {
		fmt.Println("10 é menor 100")
	}

}