//This a simples example Hello World
package main

import (
	"fmt"
)
/*
- Estrutura básica:
    - package main.
    - func main: é aqui que tudo começa, é aqui que tudo acaba.
    - import.
- Packages:
    - Pacotes são coleções de funções pré-prontas (ou não) que você pode utilizar.
    - Notação: pacote.Identificador. Exemplo: fmt.Println()
    - Documentação: fmt.Println.
- Variáveis: "uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão."
- Variáveis não utilizadas? Não pode: _ nelas.
- ...funções variádicas.
- Lição principal: package main, func main, pacote.Identificador.
 */
func mainX1() {

	//numberBytes, errors := fmt.Println("Hello World - GoLang")
	_, errors := fmt.Println("Golang RLSP")

	fmt.Println("Errors above ? " ,errors )

	x:= 10 //number
	y:= "strings" //text
	z:= true //boolean
	fmt.Println(x,y,z)
}