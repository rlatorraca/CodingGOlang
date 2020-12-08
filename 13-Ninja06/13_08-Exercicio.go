package main

import "fmt"

/*
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
 */



func main() {

	retornaumafunc()()

	/*
	x:=retornaumafunc()
	x()
	 */
}

func retornaumafunc() func() {
	return func() {
		fmt.Println("Olha eu aqui o RETORNO DA FUNCAO!")
	}
}