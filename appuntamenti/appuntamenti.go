package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Appuntamento struct {
	giorno, oraInizio, oraFine int
}

func NuovoAppuntamento(gg, h1, h2 int) (Appuntamento, bool) {
	if gg < 1 || gg > 366 || h1 < 0 || h1 > 24 || h2 < 0 || h2 > 24 || h1 > h2 {
		return Appuntamento{}, false
	}
	return Appuntamento{giorno: gg, oraInizio: h1, oraFine: h2}, true
}

func CheckSovrapposizioneAppuntamenti(app1, app2 Appuntamento) bool {
	if app1.giorno == app2.giorno && ((app1.oraInizio <= app2.oraInizio && app1.oraFine > app2.oraInizio) || (app2.oraInizio <= app1.oraInizio && app2.oraFine > app1.oraInizio)) {
		return true
	}
	return false
}

func AggiungiAppuntamento(app Appuntamento, agenda *[]Appuntamento) bool {
	for i := 0; i < len(*agenda)-1; i++ {
		if CheckSovrapposizioneAppuntamenti(app, (*agenda)[i]) {
			return false
		}
	}
	*agenda = append(*agenda, app)
	return true
}

func AppuntamentiGiornata(gg int, agenda []Appuntamento) []Appuntamento {
	var appuntamentiDelGiorno []Appuntamento
	for i := 0; i < len(agenda)-1; i++ {
		if agenda[i].giorno == gg {
			appuntamentiDelGiorno = append(appuntamentiDelGiorno, agenda[i])
		}
	}
	return appuntamentiDelGiorno
}

func main() {
	agenda := make([]Appuntamento, 0)
	var gg, h1, h2 int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		s := scanner.Text()
		fmt.Sscanf(s, "%d %d %d", &gg, &h1, &h2)
		app, _ := NuovoAppuntamento(gg, h1, h2)
		AggiungiAppuntamento(app, &agenda)
	}
	fmt.Println(agenda[:len(agenda)-2])
}
