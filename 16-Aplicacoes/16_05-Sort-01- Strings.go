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

	ss := []string{"abóbora", "maçã", "laranja", "beringela", "berinjela"}

	fmt.Println(ss)

	fmt.Println(sort.StringsAreSorted(ss))
	sort.Strings(ss)
	fmt.Println(sort.StringsAreSorted(ss))

	fmt.Println(ss)

}
