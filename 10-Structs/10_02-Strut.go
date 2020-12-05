package main

import "fmt"

/*
- Struct é um tipo de dados composto que nos permite armazenar valores de tipos *diferentes.*
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.
 */

type person struct {
	name  string
	surname string
	idade   int
}

type personProfissional struct {
	person
	title string
	salary int
}

func main() {
	person1 := person{
		name:      "João",
		surname:  "da Silva",
		idade:   22,
	}

	person2 := personProfissional{
		person : person {
			name: "Marcelo",
			surname: " da Assuncao",
			idade: 43,
		},
		title: "Software Designer",
		salary: 18000,
	}


	fmt.Println(person1)
	fmt.Println(person2)

}