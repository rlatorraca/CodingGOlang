package main

import "fmt"

/*
- x[:], x[a:], x[:b], x[a:b]
	X[ inicio:fim ]
	X[ inicio: ] // vai do 'inicio'' até o último item
	X[ : fim ] // vai do primeiro até o item na posição 'fim'
Lembrando que trabalha com posição, não valores do item

- "a" é incluso;
- "b" não é.
- Exemplo: cabeça magnética de um disco rígido (relógio, fita).
    - Off-by-one error.
- Go Playground: https://play.golang.org/p/i5ZOLKb3Fi
- É fatiando que se deleta um item de uma slice. Na prática:
    - x := append(x[:i], x[:i]...)
    - Go Playground: https://play.golang.org/p/xK2HwCqvwd
- Exercício: tente acessar todos os itens de uma slice *sem* utilizar range.
- Solução: https://play.golang.org/p/aUC9qVCobH
*/


func main() {

	//			              0           1           2             3               4
	sabores := []string{"pepperoni", "mozzarela", "hawaiana", "quatroqueijos", "marguerita"}

	fatia1 := sabores[0:2]
	fatia2 := sabores[2:4]
	fatia3 := sabores[2:]
	fatia4 := sabores[:2]
	fatia5 := sabores[2:4]

	fmt.Println(fatia1)
	fmt.Println(fatia2)
	fmt.Println(fatia3)
	fmt.Println(fatia4)
	fmt.Println(fatia5)


	fmt.Println()
	//			           0.        1.      2.        3.       4.       5.
	frutas := []string{"abacaxi", "manga", "mamao", "ameixa", "limao", "pêra"}

	todasFrutas := frutas[:]

	fmt.Println(todasFrutas)

	fmt.Println(todasFrutas[0], todasFrutas[1], todasFrutas[2], "\n")

	for i := 0; i < len(todasFrutas); i++ {

		fmt.Println(todasFrutas[i])
	}

	//Deletando
	//			              0           1           2             3               4
	tops := []string{"pepperoni", "mozzarela", "hawaiana", "quatroqueijos", "marguerita"}
	fmt.Println(tops)
	tops = append(tops[:2], tops[3:]...)  // Deleta a posicao 2 (hawaiana)

	fmt.Println(tops)
}