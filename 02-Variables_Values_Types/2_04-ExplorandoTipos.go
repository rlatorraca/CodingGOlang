package main

import (
	"fmt"
)

var novotexto = "Tem package level scope"

func main() {
	fmt.Printf("texto : %v, %T\n", textinho, textinho)
	x:= 10 //integer
	qualquerCoisa2(x, textinho)
}

func qualquerCoisa2 (x int, y string){
	fmt.Printf("Foi passado por parâmetro : %v + %v", x, y)

}