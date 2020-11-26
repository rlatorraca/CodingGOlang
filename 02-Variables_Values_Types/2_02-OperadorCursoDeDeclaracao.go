//This a simples example Hello World
package main

import (
	"fmt"
)

var texto = "Tem package level scope"

func mainX2() {

	/* := é chamado de Operador Curto de Declaracao, que parece uma MARMOTA (Gopher)
	    = So funciona dentro do Bloco de codigo (nao pode ser "package level scope")
		- serve para declarar uma variável (nome, tipo)
		- tem tipagem Automatica, com base no valor apresentado
		- deve declarar pelo menos uma variavel
		:= (declaracao + atribuição) X = (atribuicao de valor)
		%v = valor tal
		%T = tipo tal


	*/
	fmt.Printf("texto : %v, %T\n", texto, texto)

	x:= 10 //integer
	y:= "strings" //text
	z:= true //boolean
	w:= 12.2 // float
	fmt.Printf("x : %v, %T\n", x, x)
	fmt.Printf("y : %v, %T\n", y, y)
	fmt.Printf("z : %v, %T\n", z, z)
	fmt.Printf("w : %v, %T\n", w, w)

	w = 5;
	fmt.Printf("w : %v, %T\n", w, w)

	// - deve declarar pelo menos uma variavel
	w , a := 20, 22
	fmt.Printf("w : %v, %T\n", w, w)
	fmt.Printf("a : %v, %T\n", a, a)

	resultado := 10+10
	fmt.Printf("resultado : %v, %T\n", resultado, resultado)

	verdade := 10 == 10
	fmt.Printf("resultado : %v, %T\n", verdade, verdade)
}
