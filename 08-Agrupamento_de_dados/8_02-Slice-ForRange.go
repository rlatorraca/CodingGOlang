package main

import "fmt"

/*
Sliice ==> nao tem limite de tamanhos (diferentemente do Array que tem um tamanho fixo declarado)

- O que são tipos de dados compostos? Sao estrutura de dados formados por tipos primitivos como o SLICE (int, string, book)
- Wikipedia: Composite_data_type
- Effective Go: Composite literals
- ref/spec: Composite literals
- Uma slice agrupa valores de um único tipo.
- Criando uma slice: literal composta → x := []type{values}
*/


func main(){
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)


	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])
	slice[20] = 1 // Nao pode ser feito no SLICE, pois como o SLICE é feito de arrays ele deve seguir o tamanho anterior
	fmt.Println(slice[20])
}