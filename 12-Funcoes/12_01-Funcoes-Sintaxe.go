package main

import "fmt"

/*
- Qual a utilidade de funções?
    - Abstrair funcionalidade
    - Reutilização de código
- func (receiver) identifier(parameters) (returns) { code }
- A diferença entre parâmetros e argumentos:
    - Funções são definidas com parâmetros
    - Funções são chamadas com argumentos
- Em regra,  em Go é *pass by value.* ==> COPIA DO VALOR ( diferente de *pass by reference* )
    - Pass by reference, pass by copy, ... não.
	- Existe 1 excecao
- Parâmetro pode ser ...variádico.


 */
func main(){
	basica()
}

func basica() {
	fmt.Println("Oi, Brasil!")
	argumento("manhã")
	argumento("tarde")
	argumento("jfdhjf")
	valor := soma(10, 10)

	fmt.Println(valor)

	total1, quantos1, oi1 := soma2(210, 10, 1, 2, 3, 5)
	total2, quantos2, oi2 := soma3("Funcao com String + Int", 10, 1, 2, 3, 5)

	fmt.Println(total1, quantos1, oi1)
	fmt.Println(total2, quantos2, oi2)
}
func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde!")
	} else {
		fmt.Println("Oi, boa noite!")
	}
}

//Funcao com argumento (do mesmo tipo int) + retorno int
func soma(x, y int) int {
	return x + y
}

// Funcoes com multipolos retorno
func soma2(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "Resposta Finalizada!"
}

// o parametro variadido deve ser o ultimo
func soma3(s string, x ...int) (int, int, string) {
	resposta := ""
	if s == "manhã" {
		resposta = "07h - 12h"
	} else if s == "tarde" {
		resposta = "12h-18h"
	} else {
		resposta = " Nem manha e Nem tarde"
	}
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), resposta
}