package main

import "fmt"

func main() {
    var total_figurinhas_album int
    var figurinhas_possuidas int
    var vetor_geral[] int 
    var vetor_faltantes[] int
    var vetor_repetidas[] int

    fmt.Scan(&total_figurinhas_album)
    fmt.Scan(&figurinhas_possuidas)

    for i := 0; i < figurinhas_possuidas; i++ {
        var figurinha int 
        fmt.Scan(&figurinha)
        vetor_geral = append(vetor_geral, figurinha)
    }

    for i := 1; i < len(vetor_geral); i++ {
        if vetor_geral[i] == vetor_geral[i-1] {
            vetor_repetidas = append(vetor_repetidas, vetor_geral[i])
        }
    }

    for figurinha_faltante := 1; figurinha_faltante <= total_figurinhas_album; figurinha_faltante++{

        figurinha_encontrada := false

        for i := 0; i < len(vetor_geral); i++ {
            if vetor_geral[i] == figurinha_faltante {
                figurinha_encontrada = true
                break
            }
        }

        if figurinha_encontrada == false {
            vetor_faltantes = append(vetor_faltantes, figurinha_faltante)
        }
    }

    if len(vetor_repetidas) == 0 {
        fmt.Println("N")
    } else {
        for i := 0; i < len(vetor_repetidas); i++ {
            if i == len(vetor_repetidas) - 1{
                fmt.Print(vetor_repetidas[i])
            } else {
                fmt.Print(vetor_repetidas[i], " ")
            }
        }
        fmt.Println()
    }

    if len(vetor_faltantes) == 0 {
        fmt.Println("N")
    } else {
        for i := 0; i < len(vetor_faltantes); i++ {
            if i == len(vetor_faltantes) - 1{
                fmt.Print(vetor_faltantes[i])
            } else {
                fmt.Print(vetor_faltantes[i], " ")
            }
        }
        fmt.Println()
    }
        
}
