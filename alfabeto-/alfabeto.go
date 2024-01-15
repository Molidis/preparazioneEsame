package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var runes []string
	var i int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		runes = append(runes, scanner.Text())
	}

	if len(os.Args) < 2 {
		fmt.Println("Inserisci almeno una parola come argomento.")
		return
	}

	parola := strings.ToLower(os.Args[1])
	dizionario := make(map[rune]string)

	for _, lettera := range parola {
		if _, ok := dizionario[lettera]; !ok {
			fmt.Printf("%c? ", lettera)
			name := runes[i]
			i++
			dizionario[lettera] = name
		}
	}

	var translation []string
	for _, lettera := range parola {
		translation = append(translation, dizionario[lettera])
	}

	fmt.Println(strings.Join(translation, " - "))
}
