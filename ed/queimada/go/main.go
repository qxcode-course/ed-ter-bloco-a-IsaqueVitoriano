package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func getNeig(pos Pos) []Pos {
	return []Pos{
		{pos.l, pos.c - 1}, 
		{pos.l - 1, pos.c}, 
		{pos.l, pos.c + 1}, 
		{pos.l + 1, pos.c},
	}
}

func inside(grid [][]rune, pos Pos) bool {
	return !(pos.l < 0 || pos.l >= len(grid) || pos.c < 0 || pos.c >= len(grid[0]))
}

func match(grid [][]rune, pos Pos, value rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == value
}

func burnTrees(grid [][]rune, pos Pos, visitados map[Pos]bool) {
	if !match(grid, pos, '#') || visitados[pos]{
		return 
	}

	visitados[pos] = true
	grid[pos.l][pos.c] = 'o'

	for _, vizinho := range getNeig(pos) {
		burnTrees(grid, vizinho, visitados)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	visitados := make(map[Pos]bool)
	posicaoInicial := Pos{lfire, cfire}
	burnTrees(grid, posicaoInicial, visitados)
	
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
