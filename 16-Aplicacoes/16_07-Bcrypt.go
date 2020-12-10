package main

import "fmt"
import "golang.org/x/crypto/bcrypt"

/*
É uma maneira de encriptar senhas utilizando hashes.
- x/crypto/bcrypt
    - GenerateFromPassword
    - CompareHashAndPassword
- Sem Go Playground!
    - go get golang.org/x/crypto/bcrypt
 */

func main() {
	senha := "20julho1980"
	senhaerrada := "20julho1990"

	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	//func CompareHashAndPassword(hashedPassword, password []byte) error
	erro1 := bcrypt.CompareHashAndPassword(sb, []byte(senha))
	erro2 := bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada))
	if erro1 == nil {
		fmt.Println("Senha Correta")
	} else {
		fmt.Println("Senha ERRADA")
	}

	if erro2 == nil {
		fmt.Println("Senha Correta")
	} else {
		fmt.Println("Senha ERRADA")
	}


}