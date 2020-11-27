package main

import (
	"fmt"
)
/*
- Variável declarada em um code block é undefined em outro
- Para variáveis com uma abrangência maior, package level scope, utilizamos `var`
- Funciona em qualquer lugar
- Prestar atenção: chaves, colchetes, parênteses
 */
var textinho = "Tem package level scope"

func mainX3() {
	fmt.Printf("texto : %v, %T\n", textinho, textinho)
	x:= 10 //integer
	qualquerCoisa(x, textinho)
}

func qualquerCoisa (x int, y string){
	fmt.Printf("Foi passado por parâmetro : %v + %v", x, y)

}