package main

import (
	"fmt"
	"os"
	"strconv"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("nessun valore in ingresso")
		return
	}

	var previousWasEven bool
	primo, err := strconv.Atoi(os.Args[1])

	if isEven(primo) {
		previousWasEven = true
	} else {
		previousWasEven = false
	}

	if err != nil {
		fmt.Println("elemento in posizione", 1, "non valido")
		return
	}

	for i, v := range os.Args[2:] {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("elemento in posizione", i+2, "non valido")
			return
		}

		if (isEven(n) && previousWasEven) || (!isEven(n) && !previousWasEven) {
			fmt.Println("elemento in posizione", i+2, "non valido")
			return
		}

		if isEven(n) {
			previousWasEven = true
		} else {
			previousWasEven = false
		}
	}

	fmt.Println("sequenza valida")

}
