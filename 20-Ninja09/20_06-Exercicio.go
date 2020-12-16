package main

import (
	"fmt"
	"runtime"
)

/*
- Crie um programa que demonstra seu OS e ARCH.
- Rode-o com os seguintes comandos:
    - go run
    - go build => cria um arquivo para ser executado ($ ./main )
    - go install => instalad o programa no Sistema (/bin)

 */

func main() {
	fmt.Println("Seu OS é:\t", runtime.GOOS)
	fmt.Println("Seu ARCH é:\t", runtime.GOARCH)
	fmt.Println("Numero de CPUé:\t", runtime.NumCPU())
}