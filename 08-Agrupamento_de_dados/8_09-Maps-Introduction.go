package main

/*
- Utiliza o formato key:value.
- E.g. nome e telefone
- Performance excelente para lookups.
- map[key]value{ key: value }
- Acesso: m[key]
- Key sem value retorna zero. Isso pode trazer problemas.
- Para verificar: comma ok idiom.
    - v, ok := m[key]
    - ok é um boolean, true/false
- Na prática: if v, ok := m[key]; ok { }
- Para adicionar um item: m[v] = value
- Maps *não tem ordem.*
 */

import (
	"fmt"
)

func main() {

	friends := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(friends)
	fmt.Println(friends["joana"])

	//Adicionando um objeto
	friends["marquinho"] = 2342342

	fmt.Println(friends)
	fmt.Println(friends["gopher"], "\n\n")


	// comma ok idiom
	if valorDoElemento, ok := friends["naoExistente"]; !ok {
		fmt.Println("não existe!")
	} else {
		fmt.Println(valorDoElemento)
	}

}