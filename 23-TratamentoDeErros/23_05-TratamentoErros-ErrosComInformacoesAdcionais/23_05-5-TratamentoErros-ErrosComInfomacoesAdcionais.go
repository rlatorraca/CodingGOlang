package main

import (
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


// 5. type + method = error interface

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a RLSP math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrtW(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrtW(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("RLSP - math redux: square root of negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
