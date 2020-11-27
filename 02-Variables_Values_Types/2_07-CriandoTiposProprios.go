package main

import "fmt"

/*
- Revisando: tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
- Tem uma história que Bill Kennedy dizia que se um dia fizesse uma tattoo, ela diria "type is life."
- Grande parte dos aspectos mais avançados de Go dependem quase que exclusivamente de tipos.
- Como fundação para estas ferramentas, vamos aprender a declarar nossos próprios tipos.
- Revisando: tipos são fixos. Uma vez declarada uma variável como de um certo tipo, isso é imutável.
- type hotdog int → var b hotdog (main hotdog)
- Uma variável de tipo hotdog não pode ser atribuida com o valor de uma variável tipo int, mesmo que este seja o tipo subjacente de hotdog.
 */
type cachorroQuente string
var comPimenta cachorroQuente

type molhos int64
var quantidadeMolhos molhos

func main(){
	comPimenta = "Cachorro quente com pimenta em nível 1"
	quantidadeMolhos = 3
	fmt.Printf("Com pimenta : %v %T", comPimenta, comPimenta)
	fmt.Println()
	fmt.Printf("Quantidade de molhos pimenta : %v %T", quantidadeMolhos, quantidadeMolhos)
}