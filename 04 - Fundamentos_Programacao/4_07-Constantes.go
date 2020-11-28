package main

import "fmt"

/*
- São valores imutáveis.
- Podem ser tipadas ou não:
    - const oi = "Bom dia"
    - const oi string = "Bom dia"
- As não tipadas só terão um tipo atribuido a elas quando forem usadas.
    - Ex. qual o tipo de 42? int? uint? float64?
    - Ou seja, é uma flexibilidade conveniente.
- Na prática: int, float, string.
    - const x = y
    - const ( x = y )
 */
const constante = "Constante"
const (
	dia = "Bom dia"
	noite string = "Bom noite"
)


const dezConst = 10 // Os tipos das constantes são definidas NO MOMENTO DO USO
var dez = 10		// Os tipos das variáveis são definidas NO MOMENTO DA ATRIBUIÇÃO
var variavel  int

var inteira int
var flutuante float64

func main(){

	inteira = dezConst
	fmt.Println(inteira)

}