package main

import (
	"fmt"
	"os"
	"strconv"
)

func NestedSquares(n int) string {
	if n < 5 {
		return ""
	}
	var result string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("too few arguments")
		return
	}
	arg := os.Args[1]
	n, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("required integer value")
		return
	}
	result := NestedSquares(n)
	fmt.Println(result)
}
