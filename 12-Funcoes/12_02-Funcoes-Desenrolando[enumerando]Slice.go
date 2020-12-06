package main

import "fmt"

/*
- Quando temos uma slice, podemos passar os elementos individuais através "deste -> ..." operador.
- Exemplos:
 Exemplos:
    - Desenrolando uma slice de ints com como argumento para a função "soma" anterior
        - Go Playground: https://play.golang.org/p/k8O3__8UDa
    - Pode-se passar *zero* ou mais valores
        - Go Playground: https://play.golang.org/p/C238I9n7Vs
    - O parâmetro variádico deve ser o parâmetro final → ref/spec#Passing_arguments_to_..._parameters
        - Go Playground: https://play.golang.org/p/8wc2TA9xH_
        - Não roda: https://play.golang.org/p/2qTAnLWfgB

 */
func main(){

	si := []int{10, 10, 1, 2, 3, 5}

	//total := somaSlice(si) ==> ERRADO, nao podemos passar um SLICE e sim os INT do slice e fazemos isso usando ,,, (ENUMERANDO / DESENROLANDO)
	total := somaSlice(si...)

	// Passando ZERO valores, quando eh variaado
	total2 := somaSlice()


	fmt.Println(total)
	fmt.Println(total2)
}

func somaSlice(x ...int) int {

	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
