package main

import "fmt"

/*
- Utilizando este código: https://play.golang.org/p/sfyu4Is3FG
- ...use um for range loop para demonstrar os valores do canal.
 */

func main() {
	q := make(chan int)
	c := genX(q)

	fmt.Println("Enviando para receiveX(c, q)")
	receiveX(c, q)

	fmt.Println("about to exit")
}

// Coloca valores no canal
func genX(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
			fmt.Println("Incluido o numero  :", i, " no c<- i")
		}
		close(c)
		q <- 0 // Apos fechar manda ZERO para fazer o quit
	}()
	return c
}

func receiveX (c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Imprimindo Numero :", v)
		case <-q:
			fmt.Println("Saindo pelo <-q")
			return
		}
	}
}

