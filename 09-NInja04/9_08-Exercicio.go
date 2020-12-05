package main

import "fmt"

/*
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
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

	for chave, valor := range sobrenomeNomeHobbies {
		fmt.Println(chave)
		for indiceSlice, valorSlice := range valor {
			fmt.Println("\t", indiceSlice, " - ", valorSlice)
		}
	}

}