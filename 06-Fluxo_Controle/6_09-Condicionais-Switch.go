package main

import "fmt"

/*
- Switch:
    - pode avaliar uma expressão
        - switch statement == case (value)
        - default switch statement == true (bool)
    - não há fall-through por padrão
    - criando fall-through
    - default
    - cases compostos
 */

func main(){
	x:=10

	switch  {
		case x < 5:
			fmt.Println("X é menor que 5")
		case x == 5:
			fmt.Println("X é igual que 5")
		case x > 5:
			fmt.Println("X é maior que 5")
	}

	managerOffice := "Rodrigo"

	switch  managerOffice {
		case  "Marcelo":
			fmt.Println("Manager é o MARCELO")
		case "Marcos":
			fmt.Println("Manager é o MARCELO")
		case "Timoteo":
			fmt.Println("Manager é o MARCELO")
		case "Rodrigo":
			fmt.Println("Rodrigo é apenas um Engenheiro de Projeto")
			fallthrough // ira LER para proxima opcao, tamebm
		default:
			fmt.Println("Warnming ==> No Managers Today")
	}

	// Cases compostos
	teamsOnOffice := "Biro"

	switch  teamsOnOffice {
	case  "Marcelo", "José", "Bernardo", "Guiuseppe":
		fmt.Println("Manager é o MARCELO - Team 01")
	case "Marcos", "Raimundo", "Roberto", "Raí":
		fmt.Println("Manager é o MARCOS -  Team 02")
	case "Timoteo", "Biro", "Samanta", "Eliane", "Marcia":
		fmt.Println("Manager é o TIMOTEO - Team 03")
	case "Rodrigo", "Ianei", "Jorge", "Marlucio", "Tiaguinho":
		fmt.Println("Team 04 - Sem gerente")
		fallthrough // ira LER para proxima opcao, tamebm
	default:
		fmt.Println("Warnming ==> No Bussiness Day")
	}
}