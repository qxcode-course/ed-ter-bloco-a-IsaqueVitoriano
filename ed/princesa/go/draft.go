package main

import "fmt"

func imprimindo_vetor_circular(idcMatador int, vetor_circular []int) {
    fmt.Print("[")
    for i, pessoa := range vetor_circular {
        fmt.Print(" ", pessoa)
            if i == idcMatador {
                fmt.Print(">")
            }
        }
    fmt.Println(" ]")
}

func rodando_vetor(idcMatador int, vetor_circular []int) {

    for len(vetor_circular) > 1 {
        imprimindo_vetor_circular(idcMatador, vetor_circular)

        idcVitima := (idcMatador + 1) % len(vetor_circular)

        vetor_circular = append(vetor_circular[:idcVitima], vetor_circular[idcVitima + 1:]...)

        idcMatador = idcVitima % len(vetor_circular)
    }

    imprimindo_vetor_circular(idcMatador, vetor_circular)
}

func main() {
    var tamanho, inicio int
    fmt.Scan(&tamanho)
    fmt.Scan(&inicio)

    vetor_circular := make([]int, tamanho)

    for i := 0; i < tamanho; i++ {
        vetor_circular[i] = i + 1
    }

    var idcMatador int = inicio - 1

    rodando_vetor(idcMatador, vetor_circular)
}
