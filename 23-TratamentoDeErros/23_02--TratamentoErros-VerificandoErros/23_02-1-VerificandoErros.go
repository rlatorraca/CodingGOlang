package main

import "fmt"

/*
Na minha religião, underscore é pecado.
    Verifique seus erros!
    (Exceção: fmt.Println)
    Na prática:
        Exemplo 0: fmt.Println
        Exemplo 1: fmt.Scan(&var)
        Exemplo 2: os.Create → strings.NewReader → io.Copy
        Exemplo 3: os.Open → io.ReadAll
 */

func main() {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)

}