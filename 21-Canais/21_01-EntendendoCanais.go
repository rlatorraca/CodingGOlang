package main

import "fmt"

/*
- Canais são o Jeito Certo® de fazer sincronização e código concorrente.
- Eles nos permitem trasmitir valores entre goroutines.
- Servem pra coordenar, sincronizar, orquestrar, e buffering.
- Na prática:
    - make(chan type, b)
- Canais bloqueiam:
    - Eles são como corredores em uma corrida de revezamento
    - Eles tem que "passar o bastão" de maneira sincronizada
    - Se um corredor tentar passar o bastão pro próximo, mas o próximo corredor não estiver lá...
    - Ou se um corredor ficar esperando receber o bastão, mas ninguem entregar...
    - ...não dá certo.
- Exemplos:
    - Poe um valor num canal e faz um print. Block.
        - Código acima com goroutine.
        - Ou com buffer. Via de regra: má idéia; é legal em certas situações, mas em geral é melhor sempre passar o bastão de maneira sincronizada.
- Interessante: ref/spec → types
- Código:
    - Block: https://play.golang.org/p/dClS7vQlYE (não roda!)
    - Go routine: https://play.golang.org/p/ZbNCwUuiPi
    - Buffer: https://play.golang.org/p/32vYvCR7qn
    - Buffer block: https://play.golang.org/p/smeW6vigAT
    - Mais buffer: https://play.golang.org/p/Pe2pcboGiA
 */


func main() {
	/* Example 01 ==> NAO RODA (pois uma GoRoutine deve entregar a informacao do canal para outro GoRoutine
		canal := make(chan int)
		canal <- 42  ERRO => o numero entrou no canal mas NINGUEM veio pega-lo
		fmt.Println(<-canal)
	 */

	/* Example 02 ==> Com GoRoutine */
	canal := make(chan int)

	go func() {
		canal <- 42
	}()

	fmt.Println(<-canal)

	fmt.Println("======================")

	/*
		VIA DE REGRA NAO UTILIZAMOS BUFFER CHANNELS, APENAS EM CASOS ESPECIAIS
	,		= > Devemos usar GoROutines e seus matches (casamento)
	*/
	/*Example 03 ==> Com 1 Buffer */
	canal2 := make(chan int, 1) // 1 => significa  tamano do BUffer
	canal2 <- 44
	fmt.Println(<-canal2)

	fmt.Println("======================")

	/*Example 03 ==> Com 1 Buffer */
	canal3 := make(chan int, 3) // 1 => significa BUffer
	canal3 <- 42
	canal3 <- 9887
	canal3 <- 323
	fmt.Println(<-canal3)
	fmt.Println(<-canal3)
	fmt.Println(<-canal3)

	fmt.Println("======================")
}