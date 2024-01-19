package main

import (
	"fmt"
	"os"
	"strconv"
)

func max(args []int) int {
	max := args[1]
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	var mappa map[int]int
	var numeri []int
	args := os.Args[1:]
	for _, v := range args {
		numero, _ := strconv.Atoi(v)
		numeri = append(numeri, numero)
	}

	max := max(numeri)
	mappa = make(map[int]int)

	for i := 0; i <= len(numeri)-1; i++ {
		mappa[i] = numeri[i]
	}

	for i := 1; i <= max; i++ {
		for j := 0; j <= len(numeri)-1; j++ {
			if max-mappa[j] >= i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}
