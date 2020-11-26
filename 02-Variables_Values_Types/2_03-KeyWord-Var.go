package main

import (
	"fmt"
)

var textinho = "Tem package level scope"

func mainX3() {
	fmt.Printf("texto : %v, %T\n", textinho, textinho)
	x:= 10 //integer
	qualquerCoisa(x, textinho)
}

func qualquerCoisa (x int, y string){
	fmt.Printf("Foi passado por parâmetro : %v + %v", x, y)

}