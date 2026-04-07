package main
import "fmt"

func eh_Primo(n, div int) bool {
    if n <= 1 {
        return false
    }

    if n == div || n == 2 {
        return true
    }

    if n % div == 0 {
        return false
    }

    return eh_Primo(n, div + 1)

}

func geraEnesimoPrimo(n, candidato_primo int) int {
    if eh_Primo(candidato_primo, 2) {
        n--

        if n == 0 {
            return candidato_primo
        }
    }

    return geraEnesimoPrimo(n, candidato_primo + 1)
}

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(geraEnesimoPrimo(n, 2))

}
