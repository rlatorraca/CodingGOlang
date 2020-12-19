package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		log.Panicln(err)
		//		panic(err)
	}
}

/*
Panicln is equivalent to Println() followed by a call to panic().
- as funcoes Defer rondam (diferente the Fatalln que encerra automaticamente)
- um funcao recovery podera trabalhar com o panic para voltar a roda o sistema
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
