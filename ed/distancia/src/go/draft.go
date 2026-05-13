package main

import "fmt"

func podeInserir(sequencia []byte, index, L int, digito byte) bool {
    for i := 1; i <= L; i++ {
        if index-i >= 0 && sequencia[index-i] == digito {
            return false
        }
    }

    for i := 1; i <= L; i++ {
        if index+i <= len(sequencia)-1 && sequencia[index+i] == digito {
            return false
        }
    }

    return true
}

func preencheSequencia(sequencia []byte, index, L int) bool {
    if index == len(sequencia) {
        return true
    }

    if sequencia[index] != '.' {
        return preencheSequencia(sequencia, index + 1, L)
    }

    for i := 0; i <= L; i++ {
        digito := byte(i + '0')

        if podeInserir(sequencia, index, L, digito) {
            sequencia[index] = digito

            if preencheSequencia(sequencia, index + 1, L) {
                return true
            }
            sequencia[index] = '.'
        }
    }

    return false
}

func main() {
    var L int
    var entrada string

    fmt.Scan(&entrada)
    fmt.Scan(&L)

    sequencia := []byte(entrada)

    if preencheSequencia(sequencia, 0, L) {
        fmt.Println(string(sequencia))
    }
}
