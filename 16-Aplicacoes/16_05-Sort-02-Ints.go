package main

import (
	"fmt"
	"sort"
)

/*
- Sort serve para ordenar slices.
    - Como faz?
    - golang.org/pkg/ → sort
    - godoc.org/sort → examples
    - Sort altera o valor original!
- Exemplo: Ints, Strings.
- Go Playground:
    - sort.Strings: https://play.golang.org/p/Rs1NVwmg7h
    - sort.Ints: https://play.golang.org/p/I2_vsHujZa
 */


func main() {

	si := []int{123, 987, 324, 876, 234, 987, 234, 76}

	fmt.Println(si)

	sort.Ints(si)

	fmt.Println(si)

}