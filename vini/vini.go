package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome  string
	anno  int
	gradi float32
	cl    int
}

func CreaBottiglia(nome string, anno int, gradi float32, cl int) (BottigliaVino, bool) {
	if nome == "" || anno <= 0 || gradi <= 0 || cl <= 0 {
		return BottigliaVino{}, false
	}
	return BottigliaVino{nome, anno, gradi, cl}, true
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	fields := strings.Split(riga, ",")
	nome := fields[0]
	anno, _ := strconv.Atoi(fields[1])
	gradi, _ := strconv.ParseFloat(fields[2], 32)
	cl, _ := strconv.Atoi(fields[3])
	return CreaBottiglia(nome, anno, float32(gradi), cl)
}

func (b BottigliaVino) String() string {
	return fmt.Sprintf("%s, %d, %.1f°, %dcl", b.nome, b.anno, b.gradi, b.cl)
}

func main() {

	var elenco []BottigliaVino
	var sumVino int

	f, _ := os.Open(os.Args[1])
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		bottiglia, valida := CreaBottigliaDaRiga(s)
		if valida {
			elenco = append(elenco, bottiglia)
		}
	}

	oldest := 3000
	oldestIndex := 0
	var gradiMax float32
	gradiMaxIndex := 0

	for i := 0; i < len(elenco)-1; i++ {
		sumVino += elenco[i].cl
		if elenco[i].anno < oldest {
			oldestIndex = i
		}
		if elenco[i].gradi > gradiMax {
			gradiMaxIndex = i
		}
	}
	for i := 0; i < len(elenco)-1; i++ {
		fmt.Println(elenco[i])
	}
	fmt.Println("n. bottiglie:", len(elenco))
	fmt.Println("bottiglia con grado max:", elenco[gradiMaxIndex])
	fmt.Println("bottiglia più vecchia:", elenco[oldestIndex])
	fmt.Println("tot vino:", sumVino)
}
