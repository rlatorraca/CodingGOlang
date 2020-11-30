package main
/*
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/
import (
	"fmt"
)

func main() {

	anoQueNasci := 1979
	anoAtual := 2020

	for {
		if anoQueNasci > anoAtual {
			break
		}
		fmt.Printf("%v\t", anoQueNasci)
		anoQueNasci++
	}
}
