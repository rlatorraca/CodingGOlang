package main

import "fmt"

/*
- f := func(p params){ ... }
- f()
 */


func main() {

	x := 12

	y := func(x int) int {
		//fmt.Println(x, "vezes 873648 é:")
		return x * 873648
	}

	fmt.Println(x, "vezes 873648 é:", y(x))

}