package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	nl := len(grid)
	nc := len(grid[0])
	if nl == 0 {
		return 0
	}

	contador := 0
	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				contador++
				dfs(grid, i, j)
			}
		}
	}

	return contador
}

func dfs(grid[][]byte, l, c int) {
	if l < 0 || c < 0 || l >= len(grid) || c >= len(grid[0]) || grid[l][c] != '1' {
		return
	}

	grid[l][c] = '0'

	dfs(grid, l+1, c)
	dfs(grid, l-1, c)
	dfs(grid, l, c+1)
	dfs(grid, l, c-1)
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
