package main
import "fmt"

func calculaCoeficienteBinomial(n, k int) int {
    if k > n {
        return 0
    }

    if k == 0 || k == n {
        return 1 
    }

    return calculaCoeficienteBinomial(n-1, k-1) + calculaCoeficienteBinomial(n-1, k)
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)

    resultado := calculaCoeficienteBinomial(n, k)
    fmt.Println(resultado) 
}
