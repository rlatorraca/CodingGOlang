package main

/*
- Documentação é uma parte extremamente importante de fazer com que software seja acessível e sustentável.
- Documentação deve ser bem escrita e correta, mas tambem fácil de escrever e manter.
- Deve ser acoplada com o código e evoluir junto com este. Quanto mais fácil for para os programadores criarem boa documentação... melhor fica pra todos os envolvidos.
- godoc:
    - Analisa código fonte em Go, incluindo comentários, e gera documentação em HTML ou texto
    - O resultado é uma documentação firmemente atrelada ao código que documenta.
    - Por exemplo, na interface web de godoc pode-se navegar da documentação à implementação de um código com apenas um clique.
    - https://blog.golang.org/godoc-documen...
- Na prática:
    - Para documentar um tipo, uma variável, uma constante, ou um pacote, escreva um comentário imediatamente antes de sua declaração, sem linhas em branco
    - Comece a frase com o nome do elemento. No caso de pacotes, a primeira linha aparece no "package list."
    - Caso esteja escrevendo bastante documentação, utilize um arquivo doc.go. Exemplo: package fmt.
- A melhor parte dessa abordagem minimalista é que é super fácil de usar. Como resultado, muita coisa em Go, incluindo toda a standard library, já segue estas convenções.
- Outro exemplo: errors package.
 */