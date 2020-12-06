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

type pessoas struct {
nome      string
sobrenome string
idade     int
}

type arquiteto struct {
	pessoas
	especialização  string
	obrasconcluidas int
}

type dentista struct {
	pessoas
	especialização   string
	dentesarrancados int
}

type gente interface {
	oibomdia()
}

func (d dentista) oibomdia() {
	fmt.Printf("Oi, bom dia! Meu nome é %v e eu já arranquei %d dentes!\n", d.nome, d.dentesarrancados)
}

func (a arquiteto) oibomdia() {
	fmt.Printf("Oi, bom dia! Meu nome é %v e eu já fiz %d obras!\n", a.nome, a.obrasconcluidas)
}

func serhumano(g gente) {
	fmt.Println("--- func Ser Humano ---")
	switch g.(type) {
	case arquiteto:
		fmt.Println(g.(arquiteto).nome, "é uma pessoa que trabalha com", g.(arquiteto).especialização)
	case dentista:
		fmt.Println(g.(dentista).nome, "é uma pessoa que trabalha com", g.(dentista).especialização)
	}
	fmt.Println("Ele diz:")
	g.oibomdia()
}

func main() {

	alfredo := arquiteto{
		pessoas: pessoas{
			nome:      "Alfredo",
			sobrenome: "da Silva",
			idade:     30,
		},
		especialização:  "galpões",
		obrasconcluidas: 10,
	}

	rogério := dentista{
		pessoas: pessoas{
			nome:      "Rogério",
			sobrenome: "Pereira",
			idade:     60,
		},
		especialização:   "tortura",
		dentesarrancados: 7835,
	}

	fmt.Println(alfredo)
	alfredo.oibomdia()
	serhumano(alfredo)

	fmt.Println("-----")

	fmt.Println(rogério)
	rogério.oibomdia()
	serhumano(rogério)
}
