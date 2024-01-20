package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("NON OK")
		return
	}

	args := os.Args[1:]

	var somma float64
	var numeri []float64

	for _, arg := range args {
		numero, _ := strconv.Atoi(arg)
		numeri = append(numeri, float64(numero))
	}

	distanza := math.Abs(numeri[0] - numeri[1])
	somma += numeri[0]

	for i := 1; i < len(args)-1; i++ {
		if math.Abs(numeri[i]-numeri[i+1]) != distanza {
			fmt.Println("NON OK")
			return
		}
		somma += numeri[i]
	}
	somma += numeri[len(args)-1]
	fmt.Println(somma)
}
