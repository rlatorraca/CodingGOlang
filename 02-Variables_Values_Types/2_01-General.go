//This a simples example Hello World
package main

import (
	"fmt"
)

func main() {

	//numberBytes, errors := fmt.Println("Hello World - GoLang")
	_, errors := fmt.Println("Golang RLSP")

	fmt.Println("Errors above ? " ,errors )

	x:= 10 //number
	y:= "strings" //text
	z:= true //boolean
	fmt.Println(x,y,z)
}