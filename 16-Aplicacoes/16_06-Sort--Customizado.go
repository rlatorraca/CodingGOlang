package main

import (
	"fmt"
	"sort"
)

/*
- O sort que eu quero não existe. Quero fazer o meu.
- Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
    - type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
- Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
- E aí posso fazer do jeito que eu quiser.
- Exemplo:
    - struct carros: nome, consumo, potencia
    - slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
    - tipo ordenarPorPotencia
    - tipo ordenarPorConsumo
    - tipo ordenarPorLucroProDonoDoPosto
 */

type carro struct{
	nome string
	potencia int
	consumo float32
}

type ordenarPorPotencia [] carro

	func (op ordenarPorPotencia) Len() int {
		return len(op)
	}

	func (op ordenarPorPotencia) Less(i, j int) bool {
		return op[i].potencia < op[j].potencia
	}

	func (op ordenarPorPotencia) Swap(i, j int) {
		op[i], op[j] = op[j], op[i]
	}

type ordenarPorConsumo []carro
	func (op ordenarPorConsumo) Len() int {
		return len(op)
	}

	func (op ordenarPorConsumo) Less(i, j int) bool {
		return op[i].consumo > op[j].consumo
	}

	func (op ordenarPorConsumo) Swap(i, j int) {
		op[i], op[j] = op[j], op[i]
	}

type ordernarPorLucroProDonoDoPosto []carro
	func (op ordernarPorLucroProDonoDoPosto) Len() int {
		return len(op)
	}

	func (op ordernarPorLucroProDonoDoPosto) Less(i, j int) bool {
		return op[i].consumo < op[j].consumo
	}

	func (op ordernarPorLucroProDonoDoPosto) Swap(i, j int) {
		op[i], op[j] = op[j], op[i]
	}




/*
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
 */


func main() {

	carros := []carro{
		carro{"Honda Civic", 200, 12.5},
		carro{"Honda CR-V", 220, 11.5},
		carro{"Toyota Etios", 100, 15.5},
		carro{"Toyota Corolla Hybrid", 167, 22.5},
		carro{"GMC Colorado", 340, 6.5},
		carro{"Ford Focus", 204, 7.5},

	}

	fmt.Println(carros)
	fmt.Println()
	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println(carros)
	fmt.Println()
	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println(carros)
	fmt.Println()
	sort.Sort(ordernarPorLucroProDonoDoPosto(carros))
	fmt.Println(carros)
}