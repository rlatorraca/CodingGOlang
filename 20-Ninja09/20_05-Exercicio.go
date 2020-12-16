package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
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
- Utilize atomic para consertar a condição de corrida do exercício #3.
- Solução: https://github.com/ellenkorbes/aprend...

 */

var waitGroupY sync.WaitGroup

const totalGoRoutinesY = 9000000

var contadorY int32

func main() {
	inicio := time.Now().Unix()
 	contadorY = 0
 	criarGoRoutinesY(totalGoRoutinesY)
	waitGroupY.Wait()
 	fmt.Println("Total de GoRoutines : ", totalGoRoutinesY)
 	fmt.Println("Total Contador : ", contadorY)
 	fmt.Println("Tempo gasto Contador : ", time.Now().Unix() - inicio, "segundos")



}

func criarGoRoutinesY (y int){
	waitGroupY.Add(y)
	for x := 0; x < y; x++ {
		go func(){
			// func AddInt32(addr *int32, delta int32) (new int32)
			// func LoadInt32(addr *int32) (val int32)
			atomic.AddInt32(&contadorY, 1)
			atomic.LoadInt32(&contadorY)
			runtime.Gosched()
			//fmt.Println(v)
			waitGroupY.Done()
		}()
	}

}