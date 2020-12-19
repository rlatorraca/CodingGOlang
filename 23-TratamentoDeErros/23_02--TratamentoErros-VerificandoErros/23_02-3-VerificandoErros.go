package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

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
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}