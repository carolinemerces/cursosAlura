package main

import "fmt"

func main() {
    palavraDigitada := lePalavra()
    fmt.Println("A palavra digitada foi", palavraDigitada)
}

func lePalavra() string {
    var palavra string
    fmt.Print("Digite uma palavra: ")
    fmt.Scan(&palavra)
    return palavra
}