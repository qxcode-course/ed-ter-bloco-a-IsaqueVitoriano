package main

import "fmt"

type Pos struct {
    l, c int
}

func getNeig(p Pos) []Pos {
    return []Pos{
        {p.l + 1, p.c},
        {p.l - 1, p.c},
        {p.l, p.c + 1},
        {p.l, p.c - 1},
    }
}

func inside(grid [][]int, p Pos) bool {
    m := len(grid)
    n := len(grid[0])

    return p.l >= 0 && p.l < m && p.c >= 0 && p.c < n
}

func orangesRotting(grid [][]int) int {
    fila := []Pos{}
    laranjasFrescas := 0
    minutos := 0

    for i := range len(grid) {
        for j := range len(grid[0]) {
            if grid[i][j] == 2 {
                fila = append(fila, Pos{i, j})
            } else if grid[i][j] == 1 {
                laranjasFrescas++
            }
        }
    }

    for len(fila) > 0 && laranjasFrescas > 0 {
        tamanhoFila := len(fila)
        minutos++

        for i := 0; i < tamanhoFila; i++ {
            atual := fila[0]
            fila = fila[1:]

            for _, vizinho := range getNeig(atual) {
                if inside(grid, vizinho) && grid[vizinho.l][vizinho.c] == 1 {
                    grid[vizinho.l][vizinho.c] = 2
                    laranjasFrescas--
                    fila = append(fila, vizinho)
                }
            }
        }
    }

    if laranjasFrescas > 0 {
        return -1
    }

    return minutos
}

func main() {
    var m, n int
    fmt.Scan(&m, &n)

    grid := make([][]int, m)
    for i := range m {
        grid[i] = make([]int, n)

        for j := range m {
            fmt.Scan(&grid[i][j])
        }
    }

    minutos := orangesRotting(grid) 
    fmt.Println(minutos)
}