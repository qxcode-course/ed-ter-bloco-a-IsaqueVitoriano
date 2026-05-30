package main

import (
	"bufio"
	"fmt"
	"os"
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

func inside(grid [][]rune, p Pos, valor rune) bool {
	nl := len(grid)
	nc := len(grid[0])

	if p.l < 0 || p.c < 0 || p.l >= nl || p.c >= nc || grid[p.l][p.c] != valor {
		return false
	}

	return true
}

func search(grid [][]rune, inicioPos, fimPos Pos) bool {
	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()

	visitados := make(map[Pos]bool)
	caminho.Push(inicioPos)

	for !caminho.IsEmpty() {
		atual := caminho.Top()

		visitados[atual] = true
		grid[atual.l][atual.c] = '.'

		if atual == fimPos {
			break
		}

		var validos []Pos

		for _, pos := range getNeig(atual) {
			if !visitados[pos] && inside(grid, pos, ' ') {
				validos = append(validos, pos)
			}
		}

		if len(validos) > 0 {
			caminho.Push(validos[0])
		} else {
			grid[atual.l][atual.c] = ' '
			becos.Push(atual)
			caminho.Pop()
		}
	}

	return !caminho.IsEmpty()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	var inicioPos, fimPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				inicioPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				fimPos = Pos{l, c}
			}
		}
	}

	search(grid, inicioPos, fimPos)

	for _, line := range grid {
		fmt.Println(string(line))
	}

}

