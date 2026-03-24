package main

import "fmt"

func main() {

	var N int
	fmt.Scan(&N)

	fila := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&fila[i])
	}

	var qtdPessoasRetiradas int
	fmt.Scan(&qtdPessoasRetiradas)

	retiradas := make(map[int]bool)
	for i := 0; i < qtdPessoasRetiradas; i++ {
		var pessoa int
		fmt.Scan(&pessoa)
		retiradas[pessoa] = true
	}

	for i, pessoa := range fila {
		if !retiradas[pessoa] {
			if i < len(fila) - 1 {
				fmt.Print(pessoa, " ")
			} else {
				fmt.Print(pessoa, " ")
			}
		}
	}

	fmt.Println()
	


}
