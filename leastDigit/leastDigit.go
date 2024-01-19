package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkInput(s string) int {

	n, err := strconv.Atoi(s)

	if err != nil || (n < -1) {
		return -2
	} else {
		return n
	}
}

func leastDigit(n int) int {
	if n == 0 {
		return 0
	} else {
		return n % 10
	}
}

func main() {

	mappa := make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		num := checkInput(input)
		digit := leastDigit(num)
		if num > -1 {
			mappa[digit]++
		} else if num == -1 {
			return
		}
	}

	// Print the map
	for key, value := range mappa {
		fmt.Printf("%d - %d\n", key, value)
	}
}
