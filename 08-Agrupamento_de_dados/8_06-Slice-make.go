package main

/*
- Slices são feitas de arrays.
- Elas são dinâmicas, podem mudar de tamanho.
- Sempre que isso acontece, um novo array é criado e os dados são copiados.
- É conveniente, mas tem um custo computacional.
- Para otimizar as coisas, podemos utilizar make.
- make([]T, len, cap)
- "The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself.

The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume."
- len(x), cap(x)
- x[n] onde n é maior que len é out of range. Use append.
- Append maior que cap modifica o array subjacente.
- pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
- Effective Go.
 */


import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 11)

	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 12)
	slice = append(slice, 13)
	slice = append(slice, 14)
	slice = append(slice, 15)
	slice = append(slice, 16)
	slice = append(slice, 17)
	slice = append(slice, 18)
	slice = append(slice, 19)
	slice = append(slice, 20)
	slice = append(slice, 21)

	fmt.Println(slice, len(slice), cap(slice))
}
