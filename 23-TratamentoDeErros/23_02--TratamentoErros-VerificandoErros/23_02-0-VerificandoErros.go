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
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}