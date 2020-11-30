package main
/*
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/
import (
	"fmt"
)

func main() {

	anoQueEuNasci := 1979
	anoBase := anoQueEuNasci
	anoAtual := 2020

	for anoQueEuNasci <= anoAtual {
		fmt.Printf("%v\t", anoQueEuNasci)
		anoQueEuNasci++
	}
	fmt.Println("Voce terá esse ano " , anoAtual - anoBase, " anos")
}
