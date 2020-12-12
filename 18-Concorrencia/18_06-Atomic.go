package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
- Agora vamos fazer a mesma coisa, mas com atomic ao invés de mutex.
    - atomic.AddInt64
    - atomic.LoadInt64

** TESTANDO se tem  Condicao de Concorrencua: $ go run -race <nome_app).go
*/


func main(){

	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())

	var contador int64
	contador = 0
	totaldegoroutines := 15000

	var waitGroup sync.WaitGroup
	waitGroup.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println("CPUs : ", runtime.NumCPU())
	//fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("Valor do Contador: ", contador)

}