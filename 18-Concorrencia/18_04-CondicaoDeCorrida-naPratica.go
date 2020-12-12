package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
- Agora vamos dar um mergulho na documentação:
    - https://golang.org/doc/effective_go.h...
    - https://pt.wikipedia.org/wiki/Multipl...
    - O que é yield? runtime.Gosched() => tempo de entre a troca de uma tarefa em UNICO processador.
- Race condition:
        *Função 1       var     Função 2*
         Lendo: 0   →   0
         Yield          0   →   Lendo: 0
         var++: 1               Yield
         Grava: 1   →   1       var++: 1
                        1   ←   Grava: 1
         Lendo: 1   ←   1
         Yield          1   →   Lendo: 1
         var++: 2               Yield
         Grava: 2   →   2       var++: 2
                        2   ←   Grava: 2
- E é por isso que vamos ver mutex (exclusao mutua=> trave de protecao) , atomic (forma os mutex, eh mais baixo nivel que mutex)
    e, por fim, channels (mas usarremos canais em GoLang).

- Aqui vamos replicar a race condition mencionada no artigo anterior.
    - time.Sleep(time.Second) vs. runtime.Gosched()
- go help → go help build → go run -race main.go
- Como resolver? Mutex.

** TESTANDO se tem  Condicao de Concorrencua: $ go run -race <nome_app).go
*/


func main(){

	fmt.Print("CPUs : ", runtime.NumCPU())
	fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())
	contador := 0

	totalDeGoRoutines := 100000

	var waitGroup sync.WaitGroup
	waitGroup.Add(totalDeGoRoutines)

	for i := 0; i < totalDeGoRoutines; i++ {
		go func(){
			v := contador
			runtime.Gosched()
			v++
			contador = v
			waitGroup.Done()
		}()
		fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())
	}

	waitGroup.Wait()
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("Valor do Contador: ", contador)
}