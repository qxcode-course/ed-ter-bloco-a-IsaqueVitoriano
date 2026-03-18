package main

import "fmt"

func rotacionaVetor (tam_vetor, rotacoes int, vetor []int) {
    rotacoes = rotacoes % tam_vetor
    vetor = append(vetor[tam_vetor - rotacoes:], vetor[:tam_vetor - rotacoes]...)

    fmt.Print("[ ")
        for i := 0; i < len(vetor); i++ {
            fmt.Print(vetor[i], " ")
        }
    fmt.Println("]")

}

func main() {
    var tam_vetor, rotacoes, elemento int
    fmt.Scan(&tam_vetor, &rotacoes)

    var vetor []int

    for i := 1; i <= tam_vetor; i++ {
        fmt.Scan(&elemento)
        vetor = append(vetor, elemento)
    }

    rotacionaVetor(tam_vetor, rotacoes, vetor)
    
}
