package main

import (
	"fmt"
)

func imprimindo_roda(roda []int, idcMatador int) {
    matador := idcMatador

    fmt.Print("[")
    for i, pessoa := range roda {
        if i == matador {
            if pessoa > 0 {
                fmt.Printf(" %d>", pessoa)
            } else {
                fmt.Printf(" <%d", pessoa)
            }
        } else {
            fmt.Print(" ", pessoa)
        }
    }
    fmt.Println(" ]")
}

func buscando_sobrevivente(roda []int, idcMatador int) {
    for len(roda) > 1 {
        imprimindo_roda(roda, idcMatador)

        matador := roda[idcMatador]
        var idcVitima int

        if matador > 0 {
            idcVitima = (idcMatador + 1) % len(roda)
        } else {
            idcVitima = (idcMatador - 1 + len(roda)) % len(roda)
        }

        roda = append(roda[:idcVitima], roda[idcVitima+1:]...)

        if idcMatador > idcVitima {
            idcMatador--
        }

        if matador > 0 {
            idcMatador = (idcMatador + 1) % len(roda)
        } else {
            idcMatador = (idcMatador - 1 + len(roda)) % len(roda)
        }
    }
    imprimindo_roda(roda, idcMatador)
}

func main() {
    var N, E, F int
    fmt.Scan(&N, &E, &F)

    roda := make([]int, N)

    for i := 0; i < N; i++ {
        roda[i] = (i + 1) * F

        F = F * -1
    }

    idcMatador := E - 1

    buscando_sobrevivente(roda, idcMatador)
}