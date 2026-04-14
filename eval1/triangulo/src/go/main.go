package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processa(vet []int) {
	if len(vet) < 1 {
		return
	}

	var soma_vet []int

	if len(soma_vet) == 1 {
		fmt.Println("[ ", soma_vet,  " ]")
	}

	for i := range vet {
		soma_vet = append(soma_vet, vet[i])
	}

	fmt.Print("[ ")
	for i := range soma_vet {
		if i < len(soma_vet) - 1 {
			fmt.Println(soma_vet[i])
		} else {
			fmt.Println(soma_vet[i], " ")
		}
	}
	fmt.Print(" ]")

	processa(soma_vet)

	// 1. defina o ponto de parada
	// 2. monte o vetor auxiliar com os resultados das somas
	// 3. chame recursivamente a função processa para o vetor auxiliar
	// 4. imprima o vetor original
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}
