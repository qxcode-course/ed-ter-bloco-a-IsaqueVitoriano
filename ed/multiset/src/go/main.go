package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type MultiSet struct {
	data      []int
	size  	  int
	capacity  int
}

func NewMultiSet (capacity int) *MultiSet {
	return &MultiSet{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *MultiSet) String() string {
	return fmt.Sprintf("[%v]", Join(ms.data[0:ms.size], ", "))
}

func (ms *MultiSet) Insert(value int) {
	inicio := 0
	fim := ms.size - 1

	for inicio <= fim {
		meio := (inicio + fim) / 2
		if ms.data[meio] <= value {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}

	if ms.capacity == ms.size {
		ms.expand()
	}

	index := inicio
	ms.insert(value, index)
}

func (ms *MultiSet) insert(value, index int) error {
	if index < 0 || index > ms.size {
		return fmt.Errorf("index out of range")
	}

	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i - 1]
	}

	ms.data[index] = value
	ms.size += 1

	return nil
}

func (ms *MultiSet) expand() {
	newCapacity := ms.capacity * 2
	
	if newCapacity == 0 {
		ms.capacity = 1
	}

	newMultiSet := make([]int, newCapacity)

	for i := 0; i < ms.size; i++ {
		newMultiSet[i] = ms.data[i]
	}

	ms.data = newMultiSet
	ms.capacity = newCapacity
}

func (ms *MultiSet) Contains(value int) bool {
	for i := range ms.data {
		if ms.data[i] == value {
			return true
		}
	}
	return false
}

func (ms *MultiSet) Erase(value int) error {
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			ms.erase(i)
			return nil
		}
	}

	return fmt.Errorf("value not found")
}

func (ms *MultiSet) erase(index int) error {
	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i + 1]
	}

	ms.size -= 1

	return nil
}

func (ms *MultiSet) Count(value int) int {
	inicio := 0
	fim := ms.size - 1

	primeiraOcorrencia := -1

	for inicio <= fim {
		meio := (inicio + fim) / 2
		if ms.data[meio] == value {
			primeiraOcorrencia = meio
			fim = meio - 1
		} else if value > ms.data[meio] {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}

	if primeiraOcorrencia == -1 {
		return 0
	}

	inicio = 0
	fim = ms.size - 1

	ultimaOcorrencia := -1
	for inicio <= fim {
		meio := (inicio + fim) / 2
		if ms.data[meio] == value {
			ultimaOcorrencia = meio
			inicio = meio + 1
		} else if value > ms.data[meio] {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}

	return (ultimaOcorrencia - primeiraOcorrencia) + 1
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			result := ms.Erase(value)
			if result != nil {
				fmt.Println(result)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			ocorrencia := ms.Count(value)
			fmt.Println(ocorrencia)
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
