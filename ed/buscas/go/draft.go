package main

import "fmt"

func matchingStrings (str_entrada []string, str_buscada []string) {
    mapa_ocorrencia := make(map[string]int)

    for _, str_de_entrada := range str_entrada {
        mapa_ocorrencia[str_de_entrada]++
    }

    for i, str_procurada := range str_buscada {
        if i == len(str_buscada) - 1 {
            fmt.Println(mapa_ocorrencia[str_procurada])
        } else {
            fmt.Print(mapa_ocorrencia[str_procurada], " ")
        }

    }

}

func main() {
    var tam_vetor_consultas int
    var tam_vetor_buscas int 
    var vetor_consultas []string
    var vetor_buscas []string 

    fmt.Scan(&tam_vetor_consultas)
    for i := 0; i < tam_vetor_consultas; i++ {
        var str_entrada string 
        fmt.Scan(&str_entrada)
        vetor_consultas = append(vetor_consultas, str_entrada)
    }

    fmt.Scan(&tam_vetor_buscas)
    for i := 0; i < tam_vetor_buscas; i++ {
        var str_buscada string 
        fmt.Scan(&str_buscada)
        vetor_buscas = append(vetor_buscas, str_buscada)
    }

    matchingStrings(vetor_consultas, vetor_buscas)
    
}
