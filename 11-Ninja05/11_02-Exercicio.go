package main

import "fmt"

/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
 */
type pessoa2 struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main(){

	// Usado o MAKE para ENCAPSULAR os Structs
	meumapa := make(map[string]pessoa2)

	// Usado a Marmota para ENCAPSULAR os Structs
	meumapa2 := map[string]pessoa2{
	 	"Renata": pessoa2{
			nome:      "Renata",
			sobrenome: "Pimentão",
			sabores:   []string{"pistache", "morango", "baunilha"}},
		"Fredao": pessoa2{"Frederico", "da Prússia",
			[]string{"sabão em pó", "pé de coelho", "feijão"}},
	 }

	meumapa["Pimentão"] = pessoa2{
		nome:      "Renata",
		sobrenome: "Pimentão",
		sabores:   []string{"pistache", "morango", "baunilha"}}

	meumapa["da Prússia"] = pessoa2{"Frederico", "da Prússia",
		[]string{"sabão em pó", "pé de coelho", "feijão"}}

	for _, valor := range meumapa {
		fmt.Println("Meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for index, valor := range valor.sabores {
			fmt.Println("\t", index, "\t", valor)
		}
		fmt.Println("\n")
	}

	fmt.Println()

	for _, valor := range meumapa2 {
		fmt.Println("Meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for index, valor := range valor.sabores {
			fmt.Println("\t", index, "\t", valor)
		}
		fmt.Println("\n")
	}

}
