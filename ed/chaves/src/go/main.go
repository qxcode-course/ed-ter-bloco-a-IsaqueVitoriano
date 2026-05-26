package main

import (
	"fmt"
)

func main() {
	chaveamento := NewQueue[rune]()

	for selecao := 'A'; selecao <= 'P'; selecao++ {
		chaveamento.Enqueue(selecao)
	}

	for chaveamento.items.Len() > 1 {
		selecao1 := chaveamento.Dequeue()
		selecao2 := chaveamento.Dequeue()

		var gols_selecao1, gols_selecao2 int
		fmt.Scan(&gols_selecao1, &gols_selecao2)

		if gols_selecao1 > gols_selecao2 {
			chaveamento.Enqueue(selecao1)
		} else {
			chaveamento.Enqueue(selecao2)
		}
	}

	fmt.Println(string(chaveamento.Dequeue()))
}
