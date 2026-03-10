package main

import "fmt"

func main() {
    var tamanho int
    fmt.Scan(&tamanho)
    var inicio int
    fmt.Scan(&inicio)

    var vetor_circular []int

    for i := 1; i <= tamanho; i++ {
        vetor_circular = append(vetor_circular, i)
    }

    var idcVitima int 
    var idcMatador int = inicio - 1

    for len(vetor_circular) > 0 {
        fmt.Print("[ ")
            for i := 0; i < len(vetor_circular); i++ {
                fmt.Print(vetor_circular[i])

                if i == idcMatador {
                    fmt.Print(">")
                }

                if i < len(vetor_circular)-1 {
                    fmt.Print(" ")
                }
            }
        fmt.Println(" ]")

        if len(vetor_circular) == 1 {
            break
        }

        idcVitima = (idcMatador + 1) % len(vetor_circular)

        vetor_circular = append(vetor_circular[:idcVitima], vetor_circular[idcVitima + 1:]...)

        idcMatador = idcVitima % len(vetor_circular)

    }

}
