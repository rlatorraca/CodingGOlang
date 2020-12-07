package main

import "fmt"

/*
- WP: "The most common application of recursion is in mathematics and computer science, where a function being defined is applied within its own definition."
- Exemplos de recursividade: Fractais, matrioscas, efeito Droste (o efeito produzido por uma imagem que aparece dentro dela própria), GNU (“GNU is Not Unix”), etc.
- No estudo de funções: é uma função que chama a ela própria.
- Exemplo: fatoriais.
    - 4! = 4 * 3 * 2 * 1 (e no zero, deu.)
    - Com recursividade. Go Playground: https://play.golang.org/p/ujsLnUhRp_
    - Com loops. Go Playground: https://play.golang.org/p/F2VsUjYVhc
 */

func main() {
	fmt.Println(fatoriais(6))
	fmt.Println(loops(6))

}

func fatoriais(x int) int {
	if x == 1 {
		return x
	}
	return x * fatoriais(x-1)
}

func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x // total = total * x
		x--
	}
	return total
}
