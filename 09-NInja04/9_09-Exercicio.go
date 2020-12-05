package main

import "fmt"

/*
- Utilizando o exercício anterior,
   adicione uma entrada ao map e demonstre o map inteiro utilizando range.
 */

func main() {

	sobrenomeNomeHobbies := map[string][]string{
		"Onze_Eleven": []string{
			"desaparecer", "ser assustadora", "raspar o cabelo",
		},
		"senna_ayrton": []string{
			"andar de jetski", "ser milionário", "falar com sotaque de paulista mano",
		},
		"pike_rob": []string{
			"criar linguagens de programação", "usar uns ternos muito loucos",
		},
	}
	sobrenomeNomeHobbies["loureiro_kiko"] = []string{"usar os trequinho no punho", "ser do Megadeth", "tacar fogo na guitarra"}

	for chave, valor := range sobrenomeNomeHobbies {
		fmt.Println(chave)
		for indiceSlice, valorSlice := range valor {
			fmt.Println("\t", indiceSlice, " - ", valorSlice)
		}
	}
}