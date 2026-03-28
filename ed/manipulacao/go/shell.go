package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var lista []int

	for _, homem := range vet {
		if homem > 0 {
			lista = append(lista, homem)
		}
	}

	return lista
}

func getCalmWomen(vet []int) []int {
	var lista []int

	for _, mulher := range vet {
		if (mulher < 0 && mulher > -10){
			lista = append(lista, mulher)
		}
	}

	return lista
}

func sortVet(vet []int) []int {

	for i := 0; i < len(vet) - 1; i++ {
		idc_menor := i

		for j := i + 1; j < len(vet); j++ {
			if vet[j] < vet[idc_menor] {
				idc_menor = j
			}
		}

		if idc_menor != i {
			vet[i], vet[idc_menor] = vet[idc_menor], vet[i]
		}
	}

	return vet
}

func sortStress(vet []int) []int {
	for i := 0; i < len(vet) - 1; i++ {
		idc_menor := i

		for j := i + 1; j < len(vet); j++ {
			peso_j := vet[j]
			if peso_j < 0 {
				peso_j = -peso_j
			}

			peso_menor := vet[idc_menor]
			if peso_menor < 0 {
				peso_menor = -peso_menor
			}

			if peso_j < peso_menor {
				idc_menor = j
			}
		}

		if idc_menor != i {
			vet[i], vet[idc_menor] = vet[idc_menor], vet[i]
		}
	}
	return vet
}

func reverse(vet []int) []int {

	vetor_invertido := make([]int, len(vet))
	
	copy(vetor_invertido, vet)

	for i, j := 0, len(vet) - 1; i < j; i, j = i + 1, j - 1 {
		vetor_invertido[i], vetor_invertido[j] = vetor_invertido[j], vetor_invertido[i]
	}
	return vetor_invertido
}

func unique(vet []int) []int {
	ocorrencia := make(map[int]int)
	ja_adicionado := make(map[int]bool)
	var unicos []int

	for _, existe := range vet {
		ocorrencia[existe]++
	}

	for _, numero := range vet {

		if ja_adicionado[numero] {
			continue
		}
		
		ja_adicionado[numero] = true

		if ocorrencia[numero] >= 1 {
			unicos = append(unicos, numero)
		}
	}

	return unicos
}

func repeated(vet []int) []int {
	frequencia := make(map[int]int)

	var repetidos []int

	for _, numero := range vet {
		if frequencia[numero] == 0 {
			frequencia[numero]++
		} else if frequencia[numero] >= 1 {
			frequencia[numero]++

			repetidos = append(repetidos, numero)
		}
	}

	return repetidos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

