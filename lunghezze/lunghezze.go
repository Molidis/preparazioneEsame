package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mappa := make(map[int]string)

	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File non accessibile")
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := scanner.Text()
		mappa[len(s)] += " " + s
	}

	for key := 1; key < len(mappa); key++ {
		if len(mappa[key]) > key {
			fmt.Printf("%d :%s\n", key, mappa[key])
		}
	}
}
