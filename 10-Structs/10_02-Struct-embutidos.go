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
	// Structs Explicitos
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

	// Structs Implicitos
	person3 := person{"Martinha", "Bitencourt", 12}
	person4 := personProfissional{person{"Marcos Paulo", "Primeiro", 64}, "Primerio Secretario da Camara dos Deputados", 234234}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person4)
	fmt.Println(person4.name)

}