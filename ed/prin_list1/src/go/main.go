package main

import (
	"fmt"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	saida := "[ "

	atual := l.root.next

	for atual != l.root {
		if atual == sword {
			saida += fmt.Sprintf("%d> ", sword.Value)
		} else {
			saida += fmt.Sprintf("%d ", atual.Value)
		}
		atual = atual.next
	}

	return saida + "]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.next == l.root {
		return l.root.next
	}
	return it.next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
