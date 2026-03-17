package main

import (
	"fmt"
)

func main() {
    var N int
    var conta_casais = 0
    var descasados[] int

    fmt.Scan(&N)

    for i := 0; i < N; i++ {
        var animal int
        fmt.Scan(&animal)
        descasados = append(descasados, N, animal)

        if descasados[i] == animal || animal < 0 {
            descasados[i] = 0
            conta_casais++
        } else {
            descasados = append(descasados, N, animal)
        }
    }

    fmt.Println(conta_casais)

}
