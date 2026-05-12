package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	nl := len(board)
	nc := len(board[0])

	if nl == 0 {
		return
	}

	for i := 0; i < nl; i++ {
		dfs(board, i, 0)
		dfs(board, i, nc-1)
	}

	for j := 0; j < nc; j++ {
		dfs(board, 0, j)
		dfs(board, nl-1, j)
	}

	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {
			if board[i][j] == 'M' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, l, c int) {
	if l < 0 || c < 0 || l >= len(board) || c >= len(board[0]) || board[l][c] != 'O' {
		return
	}

	board[l][c] = 'M'

	dfs(board, l+1, c)
	dfs(board, l-1, c)
	dfs(board, l, c+1)
	dfs(board, l, c-1)
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
