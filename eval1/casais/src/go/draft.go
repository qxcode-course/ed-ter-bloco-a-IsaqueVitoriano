package main
import "fmt"

func main() {
    var N int
    fmt.Scan(&N)

    contaCasais := 0
    var descasados []int

    for i := 0; i < N; i++ {
        var animal int
        fmt.Scan(&animal)
        descasados = append(descasados, N, animal)

        if descasados[i] == animal || animal < 0 {
            descasados[i] = 0
            contaCasais++
        } else {
            descasados = append(descasados, N, animal)
        }
    }

    fmt.Println(contaCasais)
}
