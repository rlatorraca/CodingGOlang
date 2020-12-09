package main

import (
	"encoding/json"
	"os"
)

/*
 E agora o contrário.
- Link: https://cdn.rawgit.com/GoesToEleven/g...
- JSON-to-Go
- Tags
- Marshal/unmarshal vs. encoder/decoder
    - Marshal vai pra uma variável (precisa de uma variavel para ser manipulado)
    - Encoder "vai direto" para interface (nao precisa de um variavel para ser manipulado)
 */

type person struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {
	caracter := person{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 1000000.50,
	}

	encoder := json.NewEncoder(os.Stdout) // envia para a Interface os.Stdout (para imprimir)
	encoder.Encode(caracter)

}