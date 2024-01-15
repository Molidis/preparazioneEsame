package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var primo int
	var numero int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &primo)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &numero)
		if numero > primo {
			fmt.Print("+")
		} else if numero == primo {
			fmt.Print("=")
		} else {
			fmt.Print("-")
		}

		primo = numero
		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
		}
	}
	fmt.Println()
}
