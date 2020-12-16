package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
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

*****************
Utilize mutex para consertar a condição de corrida do exercício anterior.
- Solução: https://github.com/ellenkorbes/aprend...

 */

var waitGroupX sync.WaitGroup
var multiplexadorX sync.Mutex
const totalGoRoutinesX = 9000000

var contadorX int

func main() {
	inicio := time.Now().Unix()
 	contadorX = 0
 	criarGoRoutinesX(totalGoRoutinesX)
	waitGroupX.Wait()
 	fmt.Println("Total de GoRoutines : ", totalGoRoutinesX)
 	fmt.Println("Total Contador : ", contadorX)
 	fmt.Println("Tempo gasto Contador : ", time.Now().Unix() - inicio, "segundos")



}

func criarGoRoutinesX (y int){
	waitGroupX.Add(y)
	for x := 0; x < y; x++ {
		go func(){
			multiplexadorX.Lock()
			v := contadorX
			runtime.Gosched() // => faz um Yield (buscar outra routine)
			v++
			contadorX = v
			multiplexadorX.Unlock()
			waitGroupX.Done()
		}()
	}

}