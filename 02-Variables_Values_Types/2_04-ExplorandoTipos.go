package main

import (
	"fmt"
)

var novotexto string = "Tem package level scope"
var xx int8 = 10

func mainX4() {
	fmt.Printf("texto : %v, %T\n", textinho, textinho)
	fmt.Printf("x : %v, %T\n", xx, xx)
	qualquerCoisa2(xx, textinho)
}

func qualquerCoisa2 (xx int8, y string){
	fmt.Printf("Foi passado por parâmetro : %v + %v", xx, y)

}