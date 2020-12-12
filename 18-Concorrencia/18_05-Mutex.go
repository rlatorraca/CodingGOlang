package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
- Agora vamos resolver a race condition do programa anterior utilizando mutex.
- Mutex é mutual exclusion, exclusão mútua.
- Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez, e as outras deve aguardar sua vez "na fila."
- Na prática:
    - type Mutex
        - func (m *Mutex) Lock()
        - func (m *Mutex) Unlock()
- RWMutex

** TESTANDO se tem  Condicao de Concorrencua: $ go run -race <nome_app).go
*/


func main(){

	fmt.Print("CPUs : ", runtime.NumCPU())
	fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())
	contador := 0
	//contadorX := 1

	totalDeGoRoutines := 100000

	var waitGroup sync.WaitGroup
	waitGroup.Add(totalDeGoRoutines)

	var mu sync.Mutex

	for i := 0; i < totalDeGoRoutines; i++ {
		go func(){
			mu.Lock()
			v := contador
			//z := contadorX
			runtime.Gosched()
			v++
			//z = (contadorX+2*i)
			contador = v
			//contadorX = z
			mu.Unlock()
			waitGroup.Done()
		}()
		fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())
	}

	waitGroup.Wait()
	fmt.Println("CPUs : ", runtime.NumCPU())
	//fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("Valor do Contador: ", contador)
	//fmt.Println("Valor do ContadorX: ", contadorX)
}