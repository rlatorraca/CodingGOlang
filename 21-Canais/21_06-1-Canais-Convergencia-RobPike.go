package main

/*
- Observamos convergência quando informação de vários canais é enviada
   a um número menor de canais.
- Interessante:
- Na prática, exemplos:
    - 1. Todd:
        - Canais par, ímpar, e converge.
        - Func send manda pares pra um, ímpares pro outro, depois fecha.
        - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
        - Por fim um range retira todas as informações do canal converge.
    - 2. Rob Pike (palestra Go Concurrency Patterns):
        - Func trabalho cria um canal, cria uma go func que manda dados pra esse canal, e retorna o canal. Interessante: time.Duration(rand.Intn(1e3))
        - Func converge toma dois canais, cria um canal novo, e cria duas go funcs com for infinito que passa tudo para o canal novo. Retorna o canal novo.
        - Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2)) e usamos um for para receber dados do canal var.
 */

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := converge(trabalho("Canal FOX"), trabalho("Canal CBC"))
	for x := 0; x < 160; x++ {
		fmt.Println(<-canal)
	}

}

func trabalho(msg string) chan string {
	canal := make(chan string)

	go func(msg string, canal chan string) {
		for iter := 1; ; iter++ {
			canal <- fmt.Sprintf("Função %v diz: %v", msg, iter)
			// Dorme entre 1 e 1000 milisegundos (= 1 segundo)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3))) //1e3 = 1000
		}
	}(msg, canal)

	return canal
}

func converge(canalX, canalY chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo<- <-canalX
		}
	}()
	go func() {
		for {
			novo<- <-canalY
		}
	}()
	return novo
}