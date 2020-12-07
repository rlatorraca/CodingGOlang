package main

import "fmt"

func main(){
	temp := pesaSomenteOsNumerosImpares (somatorio, []int{11,22,33,44,55,66,77,88,99,100})
	fmt.Println(temp)
}

func somatorio(x ...int) int {
	result := 0
	for _, v := range x {
		result+=v
	}
	return result
}

func pesaSomenteOsNumerosImpares(funcaoPassada func(x ...int) int, ints []int) interface{} {
	var tempSlice []int
	for _,v := range ints {
		if v % 2 != 0 {
			tempSlice = append(tempSlice, v)
		}
	}
	total:= funcaoPassada(tempSlice...)
	return total

}
