package main

import "fmt"

/*
- Switch:
	- Pode avaliar tipos
	- Pode haver uma expressão de inicialização
 */
var x interface{

}

func main(){
	x = -512345678944445477.0

	switch x.(type) {
		case int64:
			fmt.Println("É um inteiro 64")
		case int32:
			fmt.Println("É um inteiro 32")
		case int16:
			fmt.Println("É um inteiro 16")
		case int8:
			fmt.Println("É um inteiro 8")
		case int:
			fmt.Println("É um inteiro")
		case bool:
			fmt.Println("É um boolean")
		case string:
			fmt.Println("É um String")
		case float32:
			fmt.Println("É um Float32")
		case float64:
			fmt.Println("É um Float64")
		default:
			fmt.Println("Sem tipo definido")
	}

}