package main

import "fmt"

/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
 */

func main() {

	ss := [][]string{
		[]string{
			"Joao",
			"da Silva",
			"sinuca",
		},
		[]string{
			"Marcos",
			"da Villa Maria",
			"soccer",
		},
		[]string{
			"Jorge",
			"Mario de Esposito",
			"beber",
		},
	}

	for _, v := range ss {
		fmt.Println(v)
	}

	fmt.Println("\n\n")

	for _, v := range ss {
		fmt.Println(v[0]) // Mostra apenas o primeiro valor
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}
}