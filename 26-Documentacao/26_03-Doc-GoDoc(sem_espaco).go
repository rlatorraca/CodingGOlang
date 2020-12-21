package main

/*
- godoc extrai e gera documentação de programas em Go. Funciona de duas maneiras:
    - Sem o flag http é um comando normal, mostra a documentação no stdout e é isso aí. Pode conter o flag src, que mostra o código fonte.
    - Com o flag http roda um servidor web local e mostra a documentação como página web.
- Exemplo: godoc -http=:8080 → http://localhost:8080/
 */