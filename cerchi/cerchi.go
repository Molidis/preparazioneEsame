package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type Cerchio struct {
	nome         string
	x, y, raggio float64
}

func newCircle(descr string) Cerchio {
	var c Cerchio
	fmt.Sscanf(descr, "%s %f %f %f", &c.nome, &c.x, &c.y, &c.raggio)
	return c
}

func distanzaPunti(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func tocca(cerc1, cerc2 Cerchio) bool {
	return distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y) == cerc1.raggio+cerc2.raggio+1e-6
}

func maggiore(cerc1, cerc2 Cerchio) bool {
	return cerc1.raggio > cerc2.raggio
}

func main() {
	var cerc0 Cerchio

	cerc0.nome = ""
	cerc0.x = 0.0
	cerc0.y = 0.0
	cerc0.raggio = 0.0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		descr := scanner.Text()
		if descr == "" && scanner.Err() == io.EOF {
			break
		}
		c := newCircle(descr)
		fmt.Println(c)
		if tocca(cerc0, c) {
			fmt.Print("tangente")
		} else {
			fmt.Print("non tangente")
		}
		if maggiore(cerc0, c) {
			fmt.Println(", maggiore")
		} else {
			fmt.Println(", minore o uguale")
		}
	}
}
