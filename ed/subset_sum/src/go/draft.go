package main

import "fmt"

func somaSubconjunto(inteiros []int, resultado, index, soma int) bool {
    if soma == resultado {
        return true
    }

    if index >= len(inteiros) || soma > resultado {
        return false
    }

    if somaSubconjunto(inteiros, resultado, index+1, soma + inteiros[index]) {
        return true
    }

    if somaSubconjunto(inteiros, resultado, index+1, soma) {
        return true
    }

    return false
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)

    inteiros := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&inteiros[i])
    }

    index := 0
    soma := 0
    if somaSubconjunto(inteiros, k, index, soma) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }

}
