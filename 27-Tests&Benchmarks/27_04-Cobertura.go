package main

/*

- "Cobertura" em se tratando de testes se refere a quanto do seu código, percentualmente, está sendo testado. (E antes que alguem fique neurótico querendo 100%: em geral, 70-80% tá ótimo.)
- A flag -cover faz a análise de cobertura do nosso código.
- Podemos utilizar a flag -coverprofile [arquivo] para salvar a análise de cobertura em um arquivo.
- Na prática:
    - go test -cover
    - go test -coverprofile c.out
    - go tool cover -html=c.out ← abre no browser
    - go tool cover -h ← para mais detalhes

 */
