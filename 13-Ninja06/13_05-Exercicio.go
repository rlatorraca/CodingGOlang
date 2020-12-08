package main

import (
	"fmt"
	"math"
)

/*
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
 */

type quadrado struct {
	lado float64

}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64{
	result := q.lado * q.lado
	fmt.Println("A area do quadrado: ", result)
	return result
}

func (c circulo) area() float64 {
	result := 2 * math.Phi * c.raio
	fmt.Println("A area do circulo: ", result)
	return result
}

type info interface {
	area() float64
}

func medida(i info) float64 {

	return i.area()
}

func main() {
	q1 := quadrado{
		lado: 1.34,
	}

	c1 := circulo{
		raio: 2.45,
	}

	p1 := medida(q1)
	p2 := medida(c1)

	fmt.Println("Resultado P1 : " , p1)
	fmt.Println("Resultado P2 : " , p2)
}