package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkInput(s string) int {

	n, err := strconv.Atoi(s)

	if err != nil || n < -1 {
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
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		s := scanner.Text()
		n := checkInput(s)
		if n > -1 {
			mappa[leastDigit(n)]++
		} else if n == -1 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i, "-", mappa[i])
	}
}
