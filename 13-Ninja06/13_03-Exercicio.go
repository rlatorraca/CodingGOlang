package main

import "fmt"

/*
- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
*/


func main() {
	defer fmt.Println("isso vai aparecer depois")
	fmt.Println("isso vai aparecer antes")
}
