package main

import "fmt"

/*
- Slices:
    - Tamanho: len(x)
    - Índice específico: x[i] (0-based)
- Para ver todos os itens de uma slice utilizamos o loop for com range.
- Range significa alcance, faixa, extensão.
- For range: for i, v := range x {}
*/


func main(){
	slice := []string{"banana", "maçã", "jaca", "pêssego"}

	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor:", valor)
	}

	//slice[4] = "melancia"
	slice = append(slice, "melancia")

	for _, valor := range slice {
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}

	fmt.Println()
	slice2 := []int{20, 21, 22, 23, 1, 13}

	total := 0

	for _, valor := range slice2 {

		// mesma coisa que total = total + valor
		total += valor

	}

	fmt.Println("O valor total é:", total)
}