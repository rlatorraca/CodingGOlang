package main

/*
- Slices multi-dimensionais são slices que contem slices.
- São como planilhas.
- [][]type
 */

import (
	"fmt"
)

func main() {
	ss1 := [][]int{
		// Índice: 0  1  2              // Índice:
		[]int{1, 2, 3, 4, 5, 6},        // 0
		[]int{7, 8, 9, 10, 11, 12},     // 1
		[]int{13, 14, 15, 16, 17, 18},  // 2
	}
	fmt.Println(ss1[2][4])

	ss2 := [][][][]int{

		[][][]int{
			[][]int{
				[]int{1, 2, 3, 4, 5, 6},
			},
			[][]int{
				[]int{10, 20, 30, 40, 50},
			},
		},

		[][][]int{
			[][]int{
				[]int{2, 4, 6, 8, 10},
			},
			[][]int{
				[]int{3},
			},
		},

	}
	fmt.Println(ss2[1][0][0][2])
}
