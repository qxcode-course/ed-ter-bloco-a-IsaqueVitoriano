package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct {
	data []int
	size int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (s *Set) Insert(value int) {
	inicio := 0
	fim := s.size - 1

	for inicio <= fim {
		meio := (inicio + fim) / 2 
		if s.data[meio] <= value {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}

	if s.size == s.capacity {
		s.reserve(s.capacity)
	}

	index := inicio

	s.insert(value, index)
}

func (s *Set) insert(value, index int) error {
	if index < 0 || index > s.size {
		return fmt.Errorf("index out of range")
	}

	if s.binarySearch(value) != -1 {
		return nil
	}
	
	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i - 1]
	}
	
	s.data[index] = value
	s.size += 1

	return nil
}

func (s *Set) Contains(value int) bool {
	if s.binarySearch(value) != -1 {
		return true
	}

	return false
}

func (s *Set) Erase(value int) bool {
	index := s.binarySearch(value)

	if index == -1 {
		return false
	}

	s.erase(index)
	return true
}

func (s *Set) erase(index int) error {
	for i := index; i < s.size - 1; i++ {
		s.data[i] = s.data[i + 1]
	}

	s.size -= 1
	return nil
}   

func (s *Set) reserve(newCapacity int) {
	newCapacity = s.capacity * 2

	if s.capacity == 0 {
		s.capacity = 1
	}

	newSet := make([]int, newCapacity)

	for i := 0; i < s.size; i++ {
		newSet[i] = s.data[i]
	}

	s.data = newSet
	s.capacity = newCapacity
}

func (s *Set) binarySearch(value int) int {
	inicio := 0
	fim := s.size - 1

	encontrado := -1

	for inicio <= fim {
		meio := (inicio + fim) / 2
		if s.data[meio] == value {
			return meio
		} else if s.data[meio] < value {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}

	if encontrado == -1 {
		return -1
	}

	return inicio
} 

func (s *Set) String() string {
	return fmt.Sprintf("[%v]", Join(s.data[0:s.size], ", "))
}

func Join (slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}

	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice [1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			result := v.Erase(value)
			if !result {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			result := v.Contains(value)
			if result {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
