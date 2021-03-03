package main

import "fmt"

func main() {
    pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21, 40}
    for i := 0; i < len(pontosPlanningPoker); i++ {
        fmt.Println(pontosPlanningPoker[i])
    }
}