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

func main () {
	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(z int) {
		for i := 0; i < z; i++ {
			a <- i
		}
	}(x / 2)

	go func(y int) {
		for i := 0; i < y; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
			case v := <-a:
				fmt.Println("Canal A:", v)
			case v := <-b:
				fmt.Println("Canal B:", v)
		}
	}

}