package main

import "fmt"

/*
- For
    - Repetição hierárquica (nested loops)
    - Exemplos: relógio, calendário
*/


func main(){
	for hora :=0; hora <= 23; hora++ {
		fmt.Println("Hora : ",  hora)
		for minuto :=0; minuto < 60; minuto++ {
			fmt.Printf("%v\t", minuto)
			if minuto == 19 || minuto == 39 {
				fmt.Println()
			}
		}
		fmt.Println()
	}
}