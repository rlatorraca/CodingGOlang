package main

import "fmt"

/*
 Range:
    - gofunc com for loop com send e close(chan)
    - recebe com range chan
 */

func main() {
	c := make(chan int)

	go meuloop(10, c)
	prints(c)

}

// Send
func meuloop(total int, s chan<- int) {
	for i := 0; i < total; i++ {
		s <- i // Envia os numeros para o Canal
	}
	close(s) // Encerra / Fecha o canal
}

// Receiver
func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido do canal:", v)
	}
}
