package main

import (
	"errors"
	"fmt"
	"log"
)

/*
- Para que nossas funções retornem erros customizados, podemos utilizar:
    - return errors.New()
    - return fmt.Errorf() ← tem um errors.New() embutido, olha na fonte!
    - https://golang.org/pkg/builtin/#error
- “Error values in Go aren’t special, they are just values like any other, and so you have the entire language at your disposal.” - Rob Pike
- Código:
    - 1. errors.New
    - 2. var errors.New
    - 3. fmt.Errorf
    - 4. var fmt.Errorf
    - 5. type + method = error interface
 */

//2. var errors.New

var ErrNorgateMath = errors.New("RLSP math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrtX(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrtX(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return 42, nil
}

// see use of errors.New in standard library:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go