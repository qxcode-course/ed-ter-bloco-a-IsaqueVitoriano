package main

import (
	"fmt"
)

func imprimindo_vetor_circular(idcMatador int, vetor_circular []int) {
    fmt.Print("[")
    for i, pessoa := range vetor_circular {
        if i == idcMatador {
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

func rodando_vetor(idcMatador int, vetor_circular []int) {
    for len(vetor_circular) > 1 {
        imprimindo_vetor_circular(idcMatador, vetor_circular)

        var idcVitima int
        matador := vetor_circular[idcMatador]

        if matador > 0 {
            idcVitima = (idcMatador + 1) % len(vetor_circular)
        } else {
            idcVitima = (idcMatador - 1 + len(vetor_circular)) % len(vetor_circular)
        }

        vetor_circular = append(vetor_circular[:idcVitima], vetor_circular[idcVitima + 1:]...)

        if idcMatador > idcVitima {
            idcMatador--
        }

        if matador > 0 {
            idcMatador = (idcMatador + 1) % len(vetor_circular)
        } else {
            idcMatador = (idcMatador - 1 + len(vetor_circular)) % len(vetor_circular)
        }
        
    }

    imprimindo_vetor_circular(idcMatador, vetor_circular)
}

func main() {
    var N, E, F int
    fmt.Scan(&N, &E, &F)

    vetor_circular := make([]int, N)

    for i := 0; i < N; i++ {
       elemento := i + 1

       if i % 2 == 0{
            if F == 1 {
                vetor_circular[i] = elemento
            } else {
                vetor_circular[i] = -elemento
            }
       } else {
            if F == 1 {
                vetor_circular[i] = -elemento
            } else {
                vetor_circular[i] = elemento
            }
       }
    }

    idcMatador := E - 1

    rodando_vetor(idcMatador, vetor_circular)

}
