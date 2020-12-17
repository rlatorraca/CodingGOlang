package main

import "fmt"

/*
- Select é como switch, só que pra canais, e não é sequencial.
- "A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready." — https://tour.golang.org/concurrency/5
- Na prática:
    - Exemplo 1:
        - Duas go funcs enviando X/2 numeros cada uma pra um canal
        - For loop X valores, select case ←x
    - Exemplo 2:
        - Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
        - Func 2 for infinito, select: case envia pra canal, case recebe de quit
    - Exemplo 3:
        - Chans par, ímpar, quit
        - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
        - Func receive é um select entre os três canais, encerra no quit
        - Problema!
- Go Playground:
    - 1. https://play.golang.org/p/xC3e1wBxgv
    - 2. https://play.golang.org/p/_NZqhBXN-v
    - 3. https://play.golang.org/p/rK8QwsBo0H
 */

//        - Chans par, ímpar, quit
//        - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
//        - Func receive é um select entre os três canais, encerra no quit

func main() {

	par := make(chan int)
	ímpar := make(chan int)
	quit := make(chan bool)

	go sendNumbers(par, ímpar, quit)

	receiver(par, ímpar, quit)
}

func sendNumbers(par, ímpar chan int, quit chan bool) {
	total := 1000
	for i := 0; i < total; i++ {
		if i % 2 == 0 {
			par <- i
		} else {
			ímpar <- i
		}
	}
	close(par)
	close(ímpar)
	quit <- true
}

func receiver(par, ímpar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		case v := <-ímpar:
			fmt.Println("O número", v, "é ímpar.")
		case <-quit:
			return
		}
	}
}
