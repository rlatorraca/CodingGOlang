package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
- Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race
 */

var waitGroup sync.WaitGroup

const totalGoRoutines = 5000000
var contador int

func main() {
 	contador = 0
 	criarGoRoutines(totalGoRoutines)
	waitGroup.Wait()
 	fmt.Println("Total de GoRoutines : ", totalGoRoutines)
 	fmt.Println("Total Contador : ", contador)


}

func criarGoRoutines (y int){
	waitGroup.Add(y)
	for x := 0; x < y; x++ {
		go func(){
			v := contador
			runtime.Gosched() // => faz um Yield (buscar outra routine)
			v++
			contador = v
			waitGroup.Done()
		}()
	}

}