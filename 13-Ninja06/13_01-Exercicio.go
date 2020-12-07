package main

import "fmt"

/*
- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
- Solução: https://play.golang.org/p/rxJM5fgI-9

- Revisão:
    - Funções!
        - Servem para abstrair código
        - E para reutilizar código
    - A ordem das coisas é:
        - func (receiver) identifier (parameters) (returns) { code }
    - Parâmetros vs. argumentos
    - Funções variádicas
        - Múltiplos parâmetros
        - Múltiplos argumentos
    - Métodos
    - Interfaces & polimorfismo
    - Defer
        - "Deixa pra depois!"
    - Returns
        - Múltiplos returns
        - Returns com nome (blé!)
    - Funcs como expressões
        - Atribuindo uma função a uma variável
    - Callbacks
        - Passando uma função como argumento para outra função
    - Closure
        - Capturando um scope
        - Variáveis declaradas em scopes externos são visíveis em scopes internos
    - Recursividade
        - Uma função que chama a ela mesma
        - Fatoriais
*/

func main() {
	var x int
	var y string

	fmt.Println("Zero is ", retornaInteiro())
	x,y = retornaIntEString()
	fmt.Printf("My first name is %v e minha idade é %d",y, x )
}

func retornaInteiro() int {
	return 0;
}

func retornaIntEString() (int, string) {
	return 22, "Rodrigo"
}