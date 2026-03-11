package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b, c float64
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)

    p := (a + b + c) / 2
    area_triangulo := math.Sqrt(p * ((p - a) * (p - b) * (p - c)))
    fmt.Printf("%.2f\n", area_triangulo)

}
