package main

import "fmt"

/*
- Callback: passe uma função como argumento a outra função.
 */



func main() {
	chamaUmaFuncaoPorArgumento(funcaoParaArgumentoComoCAllBack)
}

func funcaoParaArgumentoComoCAllBack() {
	fmt.Println("Olha eu aqui a FUNCAO usada COMO ARGUMENTO!")
}

func chamaUmaFuncaoPorArgumento(funcaoPassada func()) {
	fmt.Println("Prestenção:")
	funcaoPassada()
}