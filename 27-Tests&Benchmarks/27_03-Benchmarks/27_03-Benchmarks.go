package main

import "fmt"

/*
- Benchmarks nos permitem testar a velocidade ou performance do nosso código.
- Na prática:
    - Arquivo: _test.go
    - BET: Testes, Exemplos e...
    - func BenchmarkFunc (b *testing.B) { ... b.N ... }
    - go test -bench . ← todos
    - go test -bench Func ← somente Func
- go help testflag

 */
func Add(x, y int) int {
	s := []int{x, y}
	return s[0] + s[1]
}

func main() {
	fmt.Println(Add(1, 2))
}
