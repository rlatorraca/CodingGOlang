package main

/*
- Range: for k, v := range map { }
- Reiterando: maps *não tem ordem* e um range usará uma ordem aleatória.
 */

import (
	"fmt"
)

func main() {

	mapa := map[int]string{
		123: "muito legal",
		12:  "menos legal um pouquinho",
		1234: "esse é massa",
		18:  "idade de ir pra festa",
	}

	fmt.Println(mapa)

	total := 0

	for key, value := range mapa {
		fmt.Printf("Key %v - Value: %v\n", key, value)

	}

	for key, _ := range mapa {
		total+=key

	}

	fmt.Println(total)

	delete(mapa, 18)

	fmt.Println()
	for key, value := range mapa {
		fmt.Printf("Key %v - Value: %v\n", key, value)

	}
}