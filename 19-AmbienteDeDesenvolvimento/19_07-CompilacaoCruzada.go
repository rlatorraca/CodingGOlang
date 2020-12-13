package main

import (
	"fmt"
	"runtime"
)

/*
Compilacao Cruzada => compilar um arquivo em SO, para outro SO
 GOOS
- GOARCH
==> `GOOS=darwin GOARCH=amd64 go build test.go`
==> `GOOS=windows GOARCH=amd64 go build test.go`

GOOS=darwin GOARCH=amd64 go build ./19_07-CompilacaoCruzada.go
GOOS=windows GOARCH=amd64 go build ./19_07-CompilacaoCruzada.go

- https://godoc.org/runtime#pkg-constants
- git push
- git clone
- go get
- Arquivos: https://github.com/ellenkorbes/aprend...
 */

func main() {
	fmt.Println("Esse é programa do exercício de compilação cruzada. " +
		"Foi compilado num linux/amd64, e agora está rodando num sistema:",
		runtime.GOARCH, runtime.GOOS)

}