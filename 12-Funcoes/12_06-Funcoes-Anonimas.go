package main

/*
- Anonymous self-executing functions → Funções anônimas auto-executáveis.
- Funciona apenas 1x
- func(p params) { ... }(p)
- Vamos ver bastante quando falarmos de goroutines.
 */

import (
	"fmt"
)

func main() {

	x := 387

	func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}(x)

	func(y ...int) {
		total:=1
		for _, v := range y {
			total *=v
		}
		fmt.Println(total)
	}(1,2,3,4,5,6,7,8,9,10)


}