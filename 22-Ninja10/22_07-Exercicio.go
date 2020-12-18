package main

import "fmt"

/*
- Crie um programa que lance 100 goroutines onde cada uma envia 100 números a um canal;
- Tire estes números do canal e demonstre-os.
 */

func main() {
	canal := make(chan int)
	z:=0
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				canal <- z*j
				z++
			}

		}()

	}
	for k := 0; k < 10000; k++ {
		fmt.Println(k, "\t", <-canal)
	}
}

