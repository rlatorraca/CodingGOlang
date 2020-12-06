package main

import "fmt"

/*
- Em Go, valores podem ter mais que um tipo.
- Uma interface permite que um valor tenha mais que um tipo.
- Declaração: keyword identifier type → type x interface
- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
- Esse tipo será o seu tipo *e também* o tipo da interface.

- Exemplos:
    - Os tipos *profissão1* e *profissão2* contem o tipo *pessoa*
    - Cada um tem seu método *oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()*
    - Implementam a interface *gente*
    - Ambos podem acessar o função *serhumano()* que chama o método *oibomdia()* de cada *gente*
    - Tambem podemos no método *serhumano()* tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }*
    - Go Playground pré-pronto: https://play.golang.org/p/VLbo_1uE-U
    https://play.golang.org/p/zGKr7cvTPF
    - Go Playground ao vivo:
    https://play.golang.org/p/njiKbTT20Cr
- Onde se utiliza?
    - Área de formas geométricas (gobyexample.com)
    - Sort
    - DB
    - Writer interface: arquivos locais, http request/response
- Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.
 */

type people struct {
	nome      string
	sobrenome string
	idade     int
}

type dentist struct {
	people
	dentesarrancados int
	salário          float64
}

type architect struct {
	people
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentist) oibomdias() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei", x.dentesarrancados, "dentes, e ouve só: Bom dia!")
}

func (x architect) oibomdias() {
	fmt.Println("Meu nome é", x.nome, "e sou apenas um simples Arquiteto Latino Americano com Loucuro Nível: '", x.tamanhodaloucura , "' que digo: Bom dia!")
}

// Interfaces => determina o conjunto de metodos necessario para implementa-las
type gentes interface {
	oibomdias()
}

// Polimorfismo => com a mesma chamada temos funcionalidades diferentes
func serreshumanos(g gentes) {

	g.oibomdias()

	switch g.(type) {
		case dentist:
			fmt.Println("Eu ganho:", g.(dentist).salário)

		case architect:
			fmt.Println("Eu construo:", g.(architect).tipodeconstrução)
	}
}

func main() {
	mrdente := dentist{
		people: people{
			nome:      "Mister Dente",
			sobrenome: "da Silva",
			idade:     50,
		},
		dentesarrancados: 8973,
		salário:          98736.06,
	}

	mrprédio := architect{
		people: people{
			nome:      "Mister Prédio",
			sobrenome: "Sobrenome",
			idade:     51,
		},
		tipodeconstrução: "Hospícios",
		tamanhodaloucura: "HARD",
	}

	serreshumanos(mrdente)
	fmt.Println("")
	serreshumanos(mrprédio)

}
