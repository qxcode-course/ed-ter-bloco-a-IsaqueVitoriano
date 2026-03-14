package main

import (
    "fmt"
    "math"
)

func main() {
    var N int
    fmt.Scan(&N)

    var A, B int
    var menor_distancia float64 = 100 
    var idc_ganhador int

    for i := 0; i < N; i++ {
        fmt.Scan(&A, &B)

        distancia := math.Abs(float64(A) - float64(B))

        if A >= 10 && B >= 10 {
            distancia = math.Abs(float64(A) - float64(B))
            if distancia < menor_distancia {
                menor_distancia = distancia
                idc_ganhador = i
            }
        }
    }

    if idc_ganhador == 0 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(idc_ganhador)
    }

}
