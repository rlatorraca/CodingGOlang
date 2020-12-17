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

func main() {
	canal := make(chan int)
	quit := make(chan int)
	// Se usarmos as 2 como go func, terimos que usar WaitGroup
	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

func recebeQuit (canal chan int, quit chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit<- 0
}

func enviaPraCanal (canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
		// Manda informacao do canal
		case canal<- qualquercoisa:
			qualquercoisa++
		// recene informacdo do canal
		case <-quit:
			return
		}
	}
}