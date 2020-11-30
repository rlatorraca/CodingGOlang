package main
/*
 Utilizando a solução anterior, adicione as opções else if e else.
*/
import (
	"fmt"
)

func main() {
	precisoDeFerias := 18
	
	if precisoDeFerias == 2 {
		fmt.Println("DANGER => Eu preciso de férias UREGENTE")
	} else if precisoDeFerias == 1{
		fmt.Println("INFO ==> Ainda consigo me virar por um tempo")
	} else {
		fmt.Println("NICE ==> To de boa")
	}
}
