package main
import "fmt"

func subindoDegraus(n int) int {
    if n < 0 {
        return 0
    }

    if n == 0 {
        return 1
    }

    subindo_1_degrau := subindoDegraus(n-1)
    subindo_3_degraus := subindoDegraus(n-3)

    return subindo_1_degrau + subindo_3_degraus
}

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(subindoDegraus(n))
}
