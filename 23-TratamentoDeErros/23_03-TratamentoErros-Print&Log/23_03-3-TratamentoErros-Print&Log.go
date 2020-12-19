package main

import (
	"fmt"
	"log"
	"os"
)

/*
- Opções:
    - fmt.Println() → stdout
    - log.Println() → timestamp + pode-se determinar onde o erro ficará logado
    - log.Fatalln() → os.Exit(1) sem defer
    - log.Panicln() → println + panic → funcões em defer rodam; dá pra usar recover
    - panic()
- Recomendação: use log.
- Código:
    - 1. fmt.Println
    - 2. log.Println
    - 3. log.SetOutput
    - 4. log.Fatalln
    - 5. log.Panicln
    - 6. panic
- panic: http://godoc.org/builtin#panic
 */

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f) // Grava demtro do arquivo os ERROS existentes

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.