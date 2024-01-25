package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "stop" {
			break
		}

		if len(s)%2 == 0 {
			fmt.Println(reverse(s))
		}
	}
}
