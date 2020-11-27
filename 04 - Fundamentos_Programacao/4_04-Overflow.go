package main

/*
- Um uint16, por exemplo, vai de 0 a 65535.
- Que acontece se a gente tentar usar 65536?
- Ou se a gente estiver em 65535 e tentar adicionar mais 1?
- Playground: https://play.golang.org/p/t7Z4m127F2t
 */
import (
	"fmt"
)

func main() {

	var i uint16 // valores permitidos entre 0 e 65535
	i = 65535 // limite
	fmt.Println(i)
	i++
	fmt.Println(i) ///Ira resetar pra ZERO ** TOMAR CUIDADO **
	i++
	fmt.Println(i)
}