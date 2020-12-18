package main

import "fmt"

/*
- Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
 */

func main() {
	c := make(chan int)
	go insertValues(c)
	imprimeValues(c)

}

func insertValues(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func imprimeValues(canal <-chan int) {
	for v := range canal {
		fmt.Println(v)
	}
}

