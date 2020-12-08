package main

import "fmt"

/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
 */


type pessoa struct {
	nome string
	sobrenome string
	idade int
}
func (p pessoa) demonstrar() {
	fmt.Println("O nome completo dessa pessoa é:", p.nome, p.sobrenome, "\nEssa pessoa tem", p.idade, "anos.")
}

func main() {

	person := pessoa {
		nome: "Onze",
		sobrenome: "Esquisita",
		idade: 13,
	}

	person.demonstrar()

}