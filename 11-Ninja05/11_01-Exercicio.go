package main

import "fmt"

/*
 Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
 */
type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main(){
	pessoa1 := pessoa{
		nome:      "Renata",
		sobrenome: "Pimental",
		sabores:   []string{"pistache", "morango", "baunilha"}}

	pessoa2 := pessoa{"Frederico", "da Prússia",
					[]string{"sabão em pó", "pé de coelho", "feijão"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}
	fmt.Println()
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for index, v := range pessoa2.sabores {
		fmt.Println("\t", index, "\t", v)
	}

}
