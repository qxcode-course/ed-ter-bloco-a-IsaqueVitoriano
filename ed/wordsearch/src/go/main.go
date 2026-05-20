package main

import (
	"bufio"
	"fmt"
	"os"
)

func searchWord(grid[][] byte, l, c, index int, word string) bool {
	m := len(grid)
	n := len(grid[0])

	if l < 0 || c < 0 || l >= m || c >= n {
		return false
	}

	if index == len(word) {
		return true
	}

	letra := grid[l][c]

	if letra != word[index] {
		return false
	}

	grid[l][c] = 'X'
	
	existe := searchWord(grid, l+1, c, index+1, word) || searchWord(grid, l-1, c, index+1, word) || 
	searchWord(grid, l, c+1, index+1, word) || searchWord(grid, l, c-1, index+1, word) 

	grid[l][c] = letra
	return existe
} 

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	m := len(grid)
	n := len(grid[0])

	index := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if searchWord(grid, i, j, index, word) {
				return true
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
