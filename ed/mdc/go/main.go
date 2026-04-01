package main

import (
	"fmt"
)

func mdc(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	var resto int

	if a > b {
		resto = a % b
	} else {
		resto = b % a
	}

	cont := 0

	for {
		if a % resto == 0 && b % resto == 0 {
			break
		} else {
			cont++
		}
	}
	return cont
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
