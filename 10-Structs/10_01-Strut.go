package main

import "fmt"

/*
- Struct é um tipo de dados composto que nos permite armazenar valores de tipos *diferentes.*
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.
 */

type client struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := client{
		nome:      "João",
		sobrenome: "da Silva",
		fumante:   false,
	}

	cliente2 := client{"Joana", "Martire", true}

	fmt.Printf("Nome Cliente: %v %v - Fumanto: %v", cliente1.nome, cliente2.sobrenome, cliente1.fumante)
	fmt.Println()
	fmt.Printf("Nome Cliente: %v %v - Fumanto: %v", cliente2.nome, cliente2.sobrenome, cliente2.fumante)

}