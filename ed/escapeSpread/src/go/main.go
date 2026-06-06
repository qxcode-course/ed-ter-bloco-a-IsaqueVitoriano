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
		{p.l + 1, p.c},
		{p.l - 1, p.c},
		{p.l, p.c + 1},
		{p.l, p.c - 1},
	}
}

func inside(grid [][]int, p Pos) bool {
	nl := len(grid)
	nc := len(grid[0])

	return p.l >= 0 && p.l < nl && p.c >= 0 && p.c < nc
}

func prevendoFogo(grid [][]int) [][]int {
	nl := len(grid)
	nc := len(grid[0])

	tempoDoFogo := make([][]int, nl)
	for i := range nl {
		tempoDoFogo[i] = make([]int, nc)
		for j := range nc {
			tempoDoFogo[i][j] = 2000000000
		}
	}

	fila := []Pos{}

	for i := range nl {
		for j := range nc {
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

func podeEscapar(grid [][]int, tempoDoFogo [][]int, tempoDeEspera int) bool {
	nl := len(grid)
	nc := len(grid[0])
	atualMinuto := tempoDeEspera

	visitados := make([][]bool, nl)
	for i := range visitados {
		visitados[i] = make([]bool, nc)
	}

	visitados[0][0] = true

	fila := []Pos{{0,0}}

	for len(fila) > 0 {
		atualMinuto++

		nivelDaFila := len(fila)
		for i := 0; i < nivelDaFila; i++ {
			atual := fila[0]
			fila = fila[1:]

			for _, vizinho := range getNeig(atual) {
				if inside(grid, vizinho) && !visitados[vizinho.l][vizinho.c] && grid[vizinho.l][vizinho.c] != 2 {
					if vizinho.l == nl - 1 && vizinho.c == nc - 1 {
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

// Não modifique a assinatura da função maximumMinutes
// Ela é a função que será chamada no LeetCode para resolver o problema
func maximumMinutes(grid [][]int) int {
	tempoDoFogo := prevendoFogo(grid)

	inicio := 0
	fim := 1000000000
	minutes := -1

	for inicio <= fim {
		meio := (inicio + fim) / 2
		if podeEscapar(grid, tempoDoFogo, meio) {
			minutes = meio
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}

	return minutes
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
