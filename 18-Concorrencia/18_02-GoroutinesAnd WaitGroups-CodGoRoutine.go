package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)


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


var wg sync.WaitGroup

func main() {

	fmt.Println("Numero de Procesadores da Maquina:" , runtime.NumCPU())
	fmt.Println("Quantas GoRoutines ANTES de wg iniciado : ", runtime.NumGoroutine())

	wg.Add(4) // Quantas GoRoutines ??

	go funcao1()
	go funcao2()
	go funcao3()
	go funcao4()

	fmt.Println("Quantas GoRoutines DEPOIS de wg iniciado : ", runtime.NumGoroutine())

	wg.Wait() // Espera todas Go rotinas terminar

}

func funcao1() {
	for i := 0; i < 100; i++ {
		fmt.Println("funcao 1:", i)
		time.Sleep(20)
	}
	wg.Done() // Finalizado para funcao1
}

func funcao2() {
	for i := 0; i < 100; i++ {
		fmt.Println("funcao 2:", i)
		time.Sleep(10)
	}
	wg.Done() // Finalizado para funcao2
}

func funcao3() {
	for i := 0; i < 100; i++ {
		fmt.Println("funcao 3:", i)
		time.Sleep(40)
	}
	wg.Done() // Finalizado para funcao1
}

func funcao4() {
	for i := 0; i < 100; i++ {
		fmt.Println("funcao 4:", i)
		time.Sleep(5)
	}
	wg.Done() // Finalizado para funcao2
}