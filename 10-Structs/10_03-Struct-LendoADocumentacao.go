package main

import "fmt"

/*
- É importante se familiarizar com a documentação da linguagem Go.
- Neste vídeo vamos ver um pouco sobre o que a documentação diz sobre structs.
- Veremos:
    - ref/spec
        - Já vimos mais da metade dos tipos em Go!
        - Struct types.
            - x, y int
            - anonymous fields
            - promoted fields
 */

type people struct {
	name  string
	surname string
	idade   int
}

type peopleProfissional struct {
	people
	title string
	salary int
}

func main() {
	person11 := people{
		name:      "João",
		surname:  "da Silva",
		idade:   22,
	}

	person22 := peopleProfissional{
		people : people {
			name: "Marcelo",
			surname: " da Assuncao",
			idade: 43,
		},
		title: "Graphic Designer",
		salary: 11000,
	}

	// Atribuicao Concisa
	person33 := people{"Mario","Militao", 44}
	person44 := peopleProfissional{people{"Lula", "da Silva", 76}, "Ploitico", 1234567}

	fmt.Println(person11)
	fmt.Println(person22)
	fmt.Println(person11.name)
	fmt.Println(person22.people.idade)
	fmt.Println(person22.idade)
	fmt.Println(person33)
	fmt.Println(person44)

}