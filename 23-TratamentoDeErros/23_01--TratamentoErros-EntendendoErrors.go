package main

/*
- Para quem já programa em outras linguagens:
    - Em Go não temos exceções. → https://golang.org/doc/faq#exceptions
    - "We believe that coupling exceptions to a control structure, as in the try-catch-finally idiom, results in convoluted code."
    - "Go's multi-value returns make it easy to report an error without overloading the return value. A canonical error type, coupled with Go's other features, makes error handling pleasant but quite different from that in other languages."
    - Aventureiros: https://blog.golang.org/error-handlin...
- É interessante criar o hábito de lidar com erros imediatamente, similar a e.g. defer close.
- package builtin, type error interface
- package errors
 */
