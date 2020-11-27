package main

import "fmt"

/*
Strings: interpreted string literals ("") vs. raw string literals (``).
    - Rune literals. ( = cada caracter da string)
    - Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte.
		ex: int, string, bool
- Format printing: documentação.
    - Grupo #1: Print → standard out
        - func Print(a ...interface{}) (n int, err error)
        - func Println(a ...interface{}) (n int, err error)
        - func Printf(format string, a ...interface{}) (n int, err error)
            - Format verbs. (%v %T)

	- Grupo #2: Print → string, pode ser usado como variável ( e nao coloca na TELA)
        - func Sprint(a ...interface{}) string
        - func Sprintf(format string, a ...interface{}) string
        - func Sprintln(a ...interface{}) string

	- Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
        - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
        - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
        - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
 */

func main(){
	stringInterpreted := "\nBom dia\nComo vai \"você\" ?\t"
	fmt.Printf("Literal String Interpretada : %v - % T", stringInterpreted, stringInterpreted)
	stringRaw := `"\nBom dia\n               Como vai \"você\"                ?\t"`
	fmt.Printf("\nRaw String Interpretada : %v - % T", stringRaw, stringRaw)

	primeira:="\nPrimeira parte"
	segunda :="Segunda parte"
	usandoFPrint:= fmt.Sprint(primeira, " ", segunda)
	fmt.Println(usandoFPrint)
}