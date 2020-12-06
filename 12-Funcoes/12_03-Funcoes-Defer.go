package main

import "fmt"

/*
- Funções são ótimas pois tornam nosso código modular. Podemos alterar partes do nosso programa sem afetar o resto!
- Uma declaração DEFER chama uma função cuja execução ocorrerá no momento em que a função da qual ela faz parte FINALIZAR (deioxando para o final).
- Essa finalização pode ocorrer devido a um return, ao fim do code block da função, ou no caso de pânico em uma goroutine correspondente.
- "Deixa pra última hora!"
- ref/spec
- Sempre usamos para fechar um arquivo após abri-lo.
 */

func main() {
	defer fmt.Println("Primeira Linha")
	defer fmt.Println("Segunda Linha")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
}