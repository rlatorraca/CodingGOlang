package main

import "fmt"

// Declaração : reservar (comprar) um endereço de memoria para receber dados
//  => Valores ZERO
//		int : 0
//      float : 0.0
//      boolean: false
//      string: ""
//      pointer, function, interface, slices, channel, maps : nil
// Inicialização : primeiro valor coloca dentro do endereço de memoria já reservado
// atribuição de valor a variáveis: valor inserido da segunda vez em diante

var i int16 //declaracao
var j int8 = 13 //declaracao + inicializacao
var inteiro int32
var flutuante float32
var boleano bool
var cadeia string

func main(){
	i =10 // inicializacao
	fmt.Printf("i : %v, %T\n", i ,i)
	fmt.Printf("j : %v, %T\n", j ,j)

	ii := 20 // declaracao + inicializacao
	fmt.Printf("ii : %v, %T\n", ii ,ii)
	ii = 22  // atribuicao
	fmt.Printf("ii : %v, %T\n", ii ,ii)
	fmt.Printf("inteiro : %v, %T\n", inteiro ,inteiro)
	fmt.Printf("flutuante : %v, %T\n", flutuante ,flutuante)
	fmt.Printf("boleano : %v, %T\n", boleano ,boleano)
	fmt.Printf("cadeia : %v, %T\n", cadeia ,cadeia)

}