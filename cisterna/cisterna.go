package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Cisterna struct {
	larghezza, lunghezza, altezza, livello float64
}

func variazione(cisterna *Cisterna, lt int) error {

	var volumeMax, volumeAttuale float64

	volumeMax = cisterna.larghezza * cisterna.lunghezza * cisterna.altezza / 1e6
	volumeAttuale = volumeMax * cisterna.livello / cisterna.altezza / 1e6

	errIn := errors.New("puoi mettere max " + strconv.Itoa(int(volumeMax-volumeAttuale)) + " litri")
	errOut := errors.New("puoi prendere max " + strconv.Itoa(int(volumeAttuale)) + " litri")
	//fmt.Errorf sarebbe stato più veloce e meno complesso, evitando concatenazione e permettendo formattazione.

	if lt > 0 {
		if lt <= int(volumeMax-volumeAttuale) && lt > 0 && cisterna.livello+cisterna.altezza*volumeAttuale/volumeMax <= cisterna.altezza {
			volumeAttuale += float64(lt)
			cisterna.livello += cisterna.altezza * volumeAttuale / volumeMax
			return nil
		} else {
			return errIn
		}
	} else {
		if int(volumeAttuale)-lt > 0 && cisterna.livello-cisterna.altezza*volumeAttuale/volumeMax >= 0 {
			volumeAttuale -= float64(lt)
			cisterna.livello -= cisterna.altezza * volumeAttuale / volumeMax
			return nil
		} else {
			return errOut
		}
	}
}

func contenuto(cisterna Cisterna) (litri int) {
	return int(cisterna.livello * cisterna.larghezza * cisterna.lunghezza / 1e6)
}

func (cisterna Cisterna) String() string {
	return fmt.Sprintf("cisterna %.2f cm x %.2f cm x %.2f cm\nlivello attuale: %.2f cm, litri %d\n", cisterna.larghezza, cisterna.lunghezza, cisterna.altezza, cisterna.livello, contenuto(cisterna))
}

func main() {

	var cisterna Cisterna
	var errFloat1, errFloat2, errFloat3 error
	errMaggioreDiZero := errors.New("tutti gli argomenti devono essere >0")
	errNumerici := errors.New("tutti gli argomenti devono essere numerici")

	if len(os.Args) < 4 {
		fmt.Println("fornire larghezza, lunghezza, altezza")
		return
	}

	cisterna.larghezza, errFloat1 = strconv.ParseFloat(os.Args[1], 64)
	cisterna.lunghezza, errFloat2 = strconv.ParseFloat(os.Args[2], 64)
	cisterna.altezza, errFloat3 = strconv.ParseFloat(os.Args[3], 64)

	if errFloat1 != nil || errFloat2 != nil || errFloat3 != nil {
		fmt.Println(errNumerici)
		return
	} else if cisterna.larghezza <= 0 || cisterna.lunghezza <= 0 || cisterna.altezza <= 0 {
		fmt.Println(errMaggioreDiZero)
		return
	} else {
		fmt.Print(cisterna.String())
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lt, errInt := strconv.Atoi(scanner.Text())
		if lt == 9999 {
			return
		} else if errInt == nil {
			err := variazione(&cisterna, lt)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(cisterna)
			}
		}
	}
}
