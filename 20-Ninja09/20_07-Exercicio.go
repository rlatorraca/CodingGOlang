package main

import (
	"fmt"
	"runtime"
)

/*
- "If you do not leave your comfort zone, you do not remember the trip" — Brian Chesky
- Faça download e instale: https://obsproject.com/
- Escolha um tópico dos que vimos até agora. Sugestões:
    - Motivos para utilizar Go
    - Instalando Go
    - Configurando as environment variables (e.g. GOPATH)
    - Hello world
    - go commands e.g. go help
    - Variáveis
    - O operador curto de declaração
    - Constantes
    - Loops
        - init, cond, post
        - break
        - continue
    - Funções
    - func (receiver) identifier(params) (returns) { code }
    - Métodos
    - Interfaces
    - Conjuntos de métodos
    - Tipos
        - Conversão?
    - Concorrência vs. paralelismo
    - Goroutines
    - WaitGroups
    - Mutex

 */

func main() {
	fmt.Println("Seu OS é:\t", runtime.GOOS)
	fmt.Println("Seu ARCH é:\t", runtime.GOARCH)
	fmt.Println("Numero de CPUé:\t", runtime.NumCPU())
}