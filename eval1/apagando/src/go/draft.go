package main
import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	numero_contem := make(map[int]bool)
	for i := 0; i < N; i++ {
		var numero int
		fmt.Scan(&numero)

		numero_contem[numero] = true
	}

	var numeros_retirados int
	fmt.Scan(&numeros_retirados)
	for i := 0; i < numeros_retirados; i++{
		var num int
		fmt.Scan(&num)

		if numero_contem[num]{
			numero_contem[num] = false
		}
	}

	for i := range numero_contem {
		fmt.Print(numero_contem[i], " ")
	}

}
