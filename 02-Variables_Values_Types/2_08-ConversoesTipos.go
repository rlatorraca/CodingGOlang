package main

import "fmt"

/*
- Conversão de tipos é o que soa.
- Em Go não se diz casting, se diz conversion.
- a = int(b)
- ref/spec#Conversions
- Fim da sessão. Parabéns! Dicas, motivação e exercícios.
 */
type cachorroQuentao int64
var brands cachorroQuentao

func main(){
	numero := 10
	brands = 3
	fmt.Printf("Numero : %v %T", numero, numero)
	fmt.Println()
	fmt.Printf("Quantidade de molhos pimenta : %v %T", brands, brands)

	numero = int(brands + 5)
	fmt.Println()
	fmt.Printf("Numero2 : %v %T", numero, numero)
	fmt.Println()
	fmt.Printf("Quantidade de molhos pimenta2 : %v %T", brands, brands)

	brands3 := cachorroQuentao(numero - 4)
	fmt.Println()
	fmt.Printf("Numero3 : %v %T", numero, numero)
	fmt.Println()
	fmt.Printf("Quantidade de molhos pimenta3 : %v %T", brands3, brands3)
}