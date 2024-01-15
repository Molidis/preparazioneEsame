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

	for _, lettera := range parola {
		if _, ok := dizionario[lettera]; !ok {
			fmt.Printf("%c? ", lettera)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanLines)
			scanner.Scan()
			name := scanner.Text()
			dizionario[lettera] = name
		}
	}

	var translation []string
	for _, lettera := range parola {
		translation = append(translation, dizionario[lettera])
	}

	fmt.Println(strings.Join(translation, " - "))
}
