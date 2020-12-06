package main

import "fmt"

func main(){

	fmt.Println("Sum 00\t\t\tSum 01\t\t\tSum 02\t\t\tSum 03\t\t\tSum 04\t\t\tSum 05\t\t\tSum 06\t\t\tSum 07\t\t\tSum 08\t\t\tSum 09\t\t\tSum 10")
	for x:=0; x <= 10; x++ {
		for y:=0; y <= 10; y++ {
			result := x+y
			fmt.Printf("%d + %d = %d\t\t", x, y, result)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Minus 00\t\tMinus 01\t\tMinus 02\t\tMinus 03\t\tMinus 04\t\tMinus 05\t\tMinus 06\t\tMinus 07\t\tMinus 08\t\tMinus 09\t\tMinus 10")
	for x:=10; x >= 0; x-- {
		for y:=0; y <= 10; y++ {
			result := x-y
			if result > -1 {
				fmt.Printf("%d - %d = %d\t\t", x, y, result)
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Times 00\t\tTimes 01\t\tTimes 02\t\tTimes 03\t\tTimes 04\t\tTimes 05\t\tTimes 06\t\tTimes 07\t\tTimes 08\t\tTimes 09\t\tTimes 10")
	for x:=0; x <= 10; x++ {
		for y:=0; y <= 10; y++ {
			result := x*y
			fmt.Printf("%d x %d = %d\t\t", x, y, result)
		}
		fmt.Println()
	}
}
