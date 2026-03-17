package main

import "fmt"

func main() {
    var qtd_gomos int
    var direcao string
    var posX int 
    var posY int 
    var rastroX[] int 
    var rastroY[] int

    fmt.Scan(&qtd_gomos)
    fmt.Scan(&direcao)

    for i := 0; i < qtd_gomos; i++ {
        fmt.Scan(&posX, &posY)
        rastroX = append(rastroX, posX)
        rastroY = append(rastroY, posY)
    }

    for i := qtd_gomos - 1; i > 0; i-- {
        rastroX[i] = rastroX[i - 1]
        rastroY[i] = rastroY[i - 1]
    }

    switch direcao {
    case "L":
        rastroX[0] = rastroX[0] - 1
    case "U" :
        rastroY[0] = rastroY[0] - 1 
    case "R" :
        rastroX[0] = rastroX[0] + 1
    case "D" :
        rastroY[0] = rastroY[0] + 1
    }

    for i := 0; i < qtd_gomos; i++ {
        fmt.Println(rastroX[i], rastroY[i])
    }

}
