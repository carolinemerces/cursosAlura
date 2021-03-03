package main

import "fmt"

func main() {
    cidade, populacao, capital := devolveCidadeEPopulacao()
    if capital {
        fmt.Println("A capital ", cidade, "tem", populacao, "habitantes")
    } else {
        fmt.Println("A cidade ", cidade, "tem", populacao, "habitantes")
    }
}

func devolveCidadeEPopulacao() (string, int, bool) {
    return "Vila Sem Nome", 4328, true
}