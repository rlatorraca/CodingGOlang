package main

/*
- Callback é passar uma função como argumento.
- Exemplo:
    - Criando uma função que toma uma função e um []int, e usa somente os números pares como argumentos para a função.
    - Go Playground:
- Desafio: Crie uma função no programa acima que utilize somente os números *ímpares.*
 */
import "fmt"

func main(){
	temp := pegaSomentOsNumerosPares(somar, []int{11,22,33,44,55,66,77,88,99,100}...)
	fmt.Println(temp)
}

func pegaSomentOsNumerosPares(funcaoPassada func(x ...int) int, ints ...int) interface{} {
	var sliceTemp []int
	for _, v := range ints {
		if v % 2 == 0 {
			sliceTemp = append(sliceTemp, v)
		}
	}
	totalSomado := funcaoPassada(sliceTemp...)
	return totalSomado
}

func somar (x ...int) int {

	n := 0
	for _, v := range x {
		n += v
	}
	return n
}
