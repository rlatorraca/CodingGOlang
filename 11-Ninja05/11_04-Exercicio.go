package main

import "fmt"

/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
 */
func main() {

	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Stroopwafel",
		sabor:   "doce",
		ondetem: []string{"na Holanda", "na casa do seu tio rico"},
		vaibemcom: map[string]string{
			"no café da manhã": "chá",
			"no almoço":        "cafezinho",
			"na janta":         "não vai bem pq comer doce a noite engorda",
		},
	}

	fmt.Println(x)
}