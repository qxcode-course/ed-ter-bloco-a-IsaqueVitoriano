package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pos struct {
	l, c int
}

func getNeig(p Pos) []Pos {
	return []Pos {
		{p.l, p.c + 1},
		{p.l, p.c - 1},
		{p.l + 1, p.c},
		{p.l - 1, p.c},
	}
}

func inside(grid [][]int, p Pos) bool {
	m := len(grid)
	n := len(grid[0])

	return p.l >= 0 && p.c >= 0 && p.l < m && p.c < n
}

func prevendoFogo(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])

	tempoDoFogo := make([][]int, m)
	for i := range m {
		tempoDoFogo[i] = make([]int, n)
		for j := range n {
			tempoDoFogo[i][j] = 2000000000
		}
	}

	fila := []Pos{}

	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				fila = append(fila, Pos{i, j})
				tempoDoFogo[i][j] = 0
			}
		}
	}

	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]

		for _, vizinho := range getNeig(atual) {
			if inside(grid, vizinho) && grid[vizinho.l][vizinho.c] != 2 {
				novoTempoDoFogo := tempoDoFogo[atual.l][atual.c] + 1

				if novoTempoDoFogo < tempoDoFogo[vizinho.l][vizinho.c] {
					tempoDoFogo[vizinho.l][vizinho.c] = novoTempoDoFogo
					fila = append(fila, vizinho)
				}
			}
		}
	}

	return tempoDoFogo
}

func escapandoDoFogo(grid [][]int, tempoDoFogo [][]int, tempoDeEspera int) bool {
	m := len(grid)
	n := len(grid[0])

	visitados := make([][]bool, m) 
	for i := range visitados {
		visitados[i] = make([]bool, n)
	}

	visitados[0][0] = true

	atualMinuto := tempoDeEspera

	fila := []Pos{{0, 0}}

	for len(fila) > 0 {
		atualMinuto++

		nivel := len(fila)

		for i := 0; i < nivel; i++ {
			atual := fila[0]
			fila = fila[1:]

			for _, vizinho := range getNeig(atual) {
				if inside(grid, vizinho) && !visitados[vizinho.l][vizinho.c] && grid[vizinho.l][vizinho.c] != 2 {
					if vizinho.l == m - 1 && vizinho.c == n - 1 {
						if atualMinuto <= tempoDoFogo[vizinho.l][vizinho.c] {
							return true
						}
					}

					if atualMinuto < tempoDoFogo[vizinho.l][vizinho.c] {
						fila = append(fila, vizinho)
						visitados[vizinho.l][vizinho.c] = true
					}
				}
			}
		}
	}

	return false
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func maximumMinutes(grid [][]int) int {
	tempoDoFogo := prevendoFogo(grid)

	inicio := 0
	fim := 1000000000
	resposta := -1

	for inicio <= fim {
		meio := (inicio + fim)/2
		if escapandoDoFogo(grid, tempoDoFogo, meio) {
			resposta = meio
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}

	return resposta
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	nl, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	nc, _ := strconv.Atoi(scanner.Text())
	grid := make([][]int, nl)
	for i := 0; i < nl; i++ {
		grid[i] = make([]int, nc)
		for j := 0; j < nc; j++ {
			scanner.Scan()
			grid[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	result := maximumMinutes(grid)
	fmt.Println(result)
}
