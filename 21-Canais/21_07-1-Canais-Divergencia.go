package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
- Divergência é o contrário de convergência :)
- Na prática, exemplos:
    - 1. Um stream vira centenas de go funcs que depois convergem.
        - Dois canais.
        - Uma func manda X números ao primeiro canal.
        - Outra func faz um range deste canal, e para cada ítem lança uma go func que poe o retorno de trabalho() no canal dois.
        - Trabalho() é um timer aleatório pra simular workload.
        - Por fim, range canal dois demonstra os valores.
    - 2. Com throttling! Ou seja, com um número máximo de go funcs.
        - Ídem acima, mas a func que lança go funcs é assim:
        - Cria X go funcs, cada uma com um range no primeiro canal que, para cada item, poe o retorno de trabalho() no canal dois.
 */

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go sendX(20, canal1)
	go otherX(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}

}

func sendX(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func otherX(canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for v := range canal1 {
		wg.Add(1)
		go func(x int) {
			canal2 <- trabalhoX(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalhoX(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e4)))
	return n
}