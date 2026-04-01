package main
import "fmt"

func divide_rec(numero int) {
    if numero == 0 {
        return 
    }

    resto := numero % 2
    divisao := numero / 2
    divide_rec(divisao)
    fmt.Println(divisao, resto)
}

func main() {
    var num int
    fmt.Scan(&num)

    divide_rec(num)
}
