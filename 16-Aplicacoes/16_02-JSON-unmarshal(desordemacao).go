package main

import (
	"encoding/json"
	"fmt"
)

/*
 E agora o contrário.
- Link: https://cdn.rawgit.com/GoesToEleven/g...
- JSON-to-Go
- Tags
- Marshal/unmarshal vs. encoder/decoder
    - Marshal vai pra uma variável
    - Encoder "vai direto"
 */

type informacoes struct {
	Nome          string  `json:"Name"`
	Sobrenome     string  `json:"Surname"`
	Idade         int     `json:"Age"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

func main() {

	sliceOfBytes := []byte(`{"Name":"James","Surname":"Bond","Age":40,"Profissao":"Agente Secreto","Contabancaria":1000000.5}`)

	var personagem informacoes
	err := json.Unmarshal(sliceOfBytes, &personagem)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(personagem)
	fmt.Println(personagem.Profissao)
	fmt.Println(personagem.Contabancaria)
}