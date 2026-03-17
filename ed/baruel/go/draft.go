package main

import "fmt"

func main() {
    var total_figurinhas_album int
    var figurinhas_possuidas int
    var vetor_geral[] int 
    var vetor_unicas[] int
    var vetor_repetidas[] int

    fmt.Scan(&total_figurinhas_album)
    fmt.Scan(&figurinhas_possuidas)

    for i := 0; i < figurinhas_possuidas; i++ {
        var figurinha int 
        fmt.Scan(&figurinha)
        vetor_geral = append(vetor_geral, figurinhas_possuidas, figurinha)
    }

    for i := 0; i < len(vetor_geral); i++ {
        if vetor_geral[i] == vetor_geral[i-1] {
            vetor_repetidas[i] = vetor_geral[i-1]
        } else {
            vetor_unicas[i] = vetor_geral[i-1]
        }
    }

    if len(vetor_repetidas) == 0 {
        fmt.Println("N")
    } else {
        fmt.Print(vetor_repetidas)
    }
        


}
