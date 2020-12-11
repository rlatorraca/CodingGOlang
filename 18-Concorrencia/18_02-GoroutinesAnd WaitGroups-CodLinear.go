package main

import "fmt"


/*
- O código abaixo é linear. Como fazer as duas funções rodarem concorrentemente?
    - https://play.golang.org/p/XP-ZMeHUk4
- Goroutines!
- O que são goroutines? São "threads."
- O que são threads? [WP](https://pt.wikipedia.org/wiki/Thread_...))
- Na prática: go func.
- Exemplo: código termina antes da go func executar.
- Ou seja, precisamos de uma maneira pra "sincronizar" isso.
- Ah, mas então... não.
- Qualé então? sync.WaitGroup:
- Um WaitGroup serve para esperar que uma coleção de goroutines termine sua execução.
    - func Add => Quantas goroutines?
    - func Done: Finalizado!
    - func Wait: Espera tudo  mundo terminar.
- Ah, mas então... sim!
- Só pra ver: runtime.NumCPU() & runtime.NumGoroutine()
*/


func main() {
	func1()
	fmt.Println("---------------------")
	func2()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
	}
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}
}