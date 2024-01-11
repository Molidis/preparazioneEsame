package main

import (
	"fmt"
)

func Abbondante(n int) bool {

	if n <= 0 {
		return false
	}

	var somma int = 0

	for i := 1; i < n; i++ {

		if n%i == 0 {
			somma += i
		}

	}

	return somma > n

}

func main() {
	var n, j int

	fmt.Scan(&n)

	for i := 1; j < n; i++ {

		if Abbondante(i) {
			fmt.Println(i)
			j++
		}

	}
}
