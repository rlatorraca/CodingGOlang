package main

import "fmt"

/*
- Ponteiros permitem compartilhar endereços de memória. Isso é útil quando:
    - Não queremos passar grandes volumes de dados pra lá e pra cá
    - Queremos mudar um valor em sua localização original (tudo em Go é pass by value!)
- Exemplos:
    - x := 0; funçãoquemudaovalordoargumentopra1(x); Print(x)
    - x := 0; funçãoquemudaovalordo*argumentopra1(&x); Print(x)
*/


func main() {

	x := 11
	y := 11
	recebeUmValor(x)
	recebeUmPonteiro(&y)

	fmt.Println("Copia do Valor : ", x)
	fmt.Println("Ponteiro :" , y)

}

func recebeUmValor(x int) {
	x++
	fmt.Println("Na função de Copia do Valor:", x)

}

func recebeUmPonteiro(x *int) {
	*x++
	fmt.Println("Na função de Ponteiro: ", *x)
}