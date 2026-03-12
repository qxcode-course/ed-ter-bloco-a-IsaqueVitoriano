package main

import "fmt"

func main() {
    var H, P, F, D int 
    fmt.Scan(&H, &P, &F, &D)
    
    for {
        F = F + D

        if F > 15 {
            F = 0
        } else if F < 0{
            F = 15
        }

        if F == P {
            fmt.Println("N")
            break
        } else  if F == H {
            fmt.Println("S")
            break
        }
    }


}
