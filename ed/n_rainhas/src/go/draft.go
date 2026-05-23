package main

import "fmt"

func checaPosicao(matriz [][]int, linhaAtual, colunaAtual, n int) bool {
    for linha := linhaAtual-1; linha >= 0; linha-- {
        if matriz[linha][colunaAtual] == 1 {
            return false
        }
    }

    coluna_esq := colunaAtual - 1
    e_linhaDeCima := linhaAtual - 1

    for e_linhaDeCima >= 0 && coluna_esq >= 0 {
        if matriz[e_linhaDeCima][coluna_esq] == 1 {
            return false
        }

        e_linhaDeCima--
        coluna_esq--
    }

    coluna_dir := colunaAtual + 1
    d_linhaDeCima := linhaAtual - 1

    for d_linhaDeCima >= 0 && coluna_dir < n {
        if matriz[d_linhaDeCima][coluna_dir] == 1 {
            return false
        }
        
        d_linhaDeCima--
        coluna_dir++
    } 

    return true
}

func n_rainhas(matriz [][]int, n, linhaAtual int) int {
    //posicionar N rainhas em um tabuleiro N X N

    if linhaAtual == n {
        return 1
    }

    contador := 0

    for i := 0; i < n; i++ {
        if checaPosicao(matriz, linhaAtual, i, n) {
            matriz[linhaAtual][i] = 1

            contador += n_rainhas(matriz, n, linhaAtual+1)

            matriz[linhaAtual][i] = 0
        }
    }

    return contador
}

func main() {
    var n int
    fmt.Scan(&n)

    matriz := make([][]int, n)

    for i := 0; i < n; i++ {
        matriz[i] = make([]int, n)
    }

    linhaAtual := 0

    fmt.Println(n_rainhas(matriz, n, linhaAtual))
}
