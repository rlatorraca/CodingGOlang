package main

/*
- golang.org/ref/spec
- Numa declaração de constantes, o identificador iota representa números sequenciais.
- Na prática.
    - iota, iota + 1, a = iota b c, reinicia em cada const, _
- Go Playground: https://play.golang.org/p/eSrwoQjuYR
 */
import (
	"fmt"
)

const (
	a = iota + 10000000
	_ //descarta valores
	c
	w
	_ //descarta valores
	z

)

const (
	const1 = iota
	const2 = iota
	_  = iota
	const4 = iota
	const5 = iota
	_ = iota
	const7 = iota
)

func main() {
	fmt.Println(a, c, w, z)
	fmt.Println(const1, const2, const4, const5, const7)
}