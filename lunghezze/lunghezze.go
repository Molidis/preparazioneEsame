package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	mappa := make(map[string]int)
	var keys []string

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
		mappa[s] = len(s)
	}

	for k := range mappa {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%d : %s\n", mappa[k], k)
	}
}
