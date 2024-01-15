package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Inserisci almeno una parola come argomento.")
		return
	}

	parola := strings.ToLower(os.Args[1])
	dizionario := make(map[rune]string)

	for _, letter := range parola {
		if _, ok := dizionario[letter]; !ok {
			fmt.Printf(" %c? ", letter)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			name := scanner.Text()
			dizionario[letter] = name
		}
	}

	var translation []string
	for _, letter := range parola {
		translation = append(translation, dizionario[letter])
	}

	fmt.Println(strings.Join(translation, " - "))
}
