package main

import "fmt"

/*
- Todos os valores ficam armazenados na memória.
- Toda localização na memória possui um endereço.
- Um pointeiro se refere a esse endereço.
- Notações:
    - &variável mostra o endereço de uma variável
        - %T: variável vs. &variável
    - *variável faz de-reference, mostra o valor que consta nesse endereço
    - ????: *&var funciona!
    - *type é um tipo que contem o endereço de um valor do tipo type, nesse caso * não é um operador
- Exemplo: a := 0; b := &a; *b++
 */

func main() {

	x := 0

	y := &x

	// *&x => mostra o VALOR do ENDERECO da variavel x
	// *x => mostra o VALOR de x
	// &x => mostra o ENDERECO de MEMORIA de x
	fmt.Print(x, y)

	*y++ // incrementa o valor existente na posicao de memoraia da variavel x

	fmt.Println(*y)
	fmt.Printf("Tipo de X : %T, tipo de Y : %T\n", x, y)
	fmt.Print(x, y)
}