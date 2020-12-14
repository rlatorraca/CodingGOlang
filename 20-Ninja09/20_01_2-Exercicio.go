package main

import (
	"fmt"
	"sync"
)

/*
- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
- Solução:
    - https://github.com/ellenkorbes/aprend...
    - https://github.com/ellenkorbes/aprend...
 */


var wg2 sync.WaitGroup

func main() {
	wg2.Add(8)
	go func() { fmt.Println("Eu sou a goroutine número: 1"); wg2.Done() }()
	go func() { fmt.Println("Eu sou a goroutine número: 2"); wg2.Done() }()
	go func() { fmt.Println("Eu sou a goroutine número: 3"); wg2.Done() }()
	go func() { fmt.Println("Eu sou a goroutine número: 4"); wg2.Done() }()
	go func() { fmt.Println("Eu sou a goroutine número: 5"); wg2.Done() }()
	go func() { fmt.Println("Eu sou a goroutine número: 6"); wg2.Done() }()
	go func() { fmt.Println("Eu sou a goroutine número: 7"); wg2.Done() }()
	go func() { fmt.Println("Eu sou a goroutine número: 8"); wg2.Done() }()
	wg2.Wait()
}