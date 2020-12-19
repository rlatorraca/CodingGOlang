package main

import "fmt"

/*
- Base-10: decimal, 0–9 ==> %d
- Base-2: binário, 0–1 ==> %b
- Base-16: hexadecimal, 0–f ==> %#x
- https://docs.google.com/document/d/1G...
 */

func main(){
	num := 128
	fmt.Printf("%d\t %b\t %#x", num, num, num)
}