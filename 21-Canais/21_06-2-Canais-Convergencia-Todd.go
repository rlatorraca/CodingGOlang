package main

import (
	"fmt"
	"sync"
)

/*
- Observamos convergência quando informação de vários canais é enviada
   a um número menor de canais.
- Interessante:
- Na prática, exemplos:
    - 1. Todd:
        - Canais par, ímpar, e converge.
        - Func send manda pares pra um, ímpares pro outro, depois fecha.
        - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
        - Por fim um range retira todas as informações do canal converge.
    - 2. Rob Pike (palestra Go Concurrency Patterns):
        - Func trabalho cria um canal, cria uma go func que manda dados pra esse canal, e retorna o canal. Interessante: time.Duration(rand.Intn(1e3))
        - Func converge toma dois canais, cria um canal novo, e cria duas go funcs com for infinito que passa tudo para o canal novo. Retorna o canal novo.
        - Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2)) e usamos um for para receber dados do canal var.
 */

func main() {
	par := make(chan int)
	ímpar := make(chan int)
	converge := make(chan int)

	go envia(par, ímpar)
	go recebe(par, ímpar, converge)

	for v := range converge {
		fmt.Println("Valor recebido:", v)
	}

}

func envia(par, impar chan int) {
	x := 100
	for n := 1; n <= x; n++ {
		if n%2 == 0 {
			par<- n
		} else {
			impar<- n
		}
	}
	close(par)
	close(impar)
}

func recebe(par, impar, converge chan int) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		for v := range par {
			converge <- v
		}
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		for v := range impar {
			converge <- v
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
	close(converge)
}

// - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
// - Por fim um range retira todas as informações do canal converge.