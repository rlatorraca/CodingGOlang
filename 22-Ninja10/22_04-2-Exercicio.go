package main

import "fmt"

/*
- Utilizando este código: https://play.golang.org/p/sfyu4Is3FG
- ...use um for range loop para demonstrar os valores do canal.
 */

func main() {
	q := make(chan int)
	c := genY(q)

	fmt.Println("Enviando para receiveX(c, q)")
	receiveY(c, q)

	fmt.Println("about to exit")
}

// Coloca valores no canal
func genY(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("Incluido o numero  :", i, " no c<- i")
		}
		close(c)
		q <- 0 // Apos fechar manda ZERO para fazer o quit
	}()
	return c
}

func receiveY (c <-chan int, q chan int) {
	for {
		select {
		case val, ok := <-c:

			if ok {
				fmt.Println("Valor vindo do c <-chan:", val)
			} else {
				return
			}
		case val, ok := <-q:

			if ok {
				fmt.Println("O canal quit enviou:", val)
			} else {
				return
			}

		}

	}

}

