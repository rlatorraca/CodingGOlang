package main

import (
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
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("ERROR : ", err)
				log.Println("ERROR : ", err) // grava data e hoario do Erro
		//		log.Fatalln(err)
		//		panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
*/

// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.