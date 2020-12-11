package main

import (
	"fmt"
	"sort"
)

/*
Partindo do código abaixo, ordene os []user por idade e sobrenome.
    - https://play.golang.org/p/BVRZTdlUZ_
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
 */

type users struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type porIdade []users

func (pi porIdade) Len() int           { return len(pi) }
func (pi porIdade) Less(i, j int) bool { return pi[i].Age < pi[j].Age }
func (pi porIdade) Swap(i, j int)      { pi[i], pi[j] = pi[j], pi[i] }

type porSobrenome []users

func (ps porSobrenome) Len() int           { return len(ps) }
func (ps porSobrenome) Less(i, j int) bool { return ps[i].Last < ps[j].Last }
func (ps porSobrenome) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

func main() {
	u1 := users{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := users{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := users{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []users{u1, u2, u3}

	for _, v := range users {
		sort.Strings(v.Sayings) // ordena o Sayings
	}

	sort.Sort(porIdade(users))
	fmt.Println("Ordenado por idade:\n")
	formatarSaida(users)

	sort.Sort(porSobrenome(users))
	fmt.Println("Ordenado por sobrenome:\n")
	formatarSaida(users)

}

func formatarSaida(x []users) {
	for i, v := range x {
		fmt.Println("Id :",i, "\tIdade:", v.Age, "\tNome completo:", v.First, v.Last, "\n")
		for _, c := range v.Sayings {
			fmt.Println("\t", c)
		}
		fmt.Println()

	}
}
