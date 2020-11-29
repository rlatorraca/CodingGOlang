package main

import "fmt"

/*
- For: inicialização, condição, pós
- For: condição ("while")
- For: ...ever? (http servers)
- For: break
- golang.org/ref/spec#For_statements, Effective Go
- (Range vem mais pra frente.)
*/

func main(){

	// NAO EXISTE WHILE em GO lang

	/*
		Adaptamos em FOR
	 */
	x := 0
	for x< 10 {
		fmt.Printf("%v\t", x)
		x++
	}

	fmt.Println()

	//Loop Infinito
	y := 1

	for {
		if y <= 20 {
			fmt.Printf("%v\t", y)
			y++
		} else {
			fmt.Println("Chegou em 20 e saiu ....\n ")
			break // Sai do Loop
		}
	}
	fmt.Println("Fora do loop ")


}
