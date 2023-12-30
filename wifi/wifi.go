package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Wifi struct {
	ssid       string
	channel    int
	frequency  int
	signal_dBm int
}

func NewWifi(ssid string, channel int, frequency int, signal_dBm int) (Wifi, bool) {
	if ssid != "" && (channel > 0 && channel < 15 && frequency > 2412 && frequency < 2484) || (channel > 7 && channel < 196 && frequency > 5035 && frequency < 5825) && ((frequency > 2412 && frequency < 2484) || (frequency > 5035 && frequency < 5980)) && signal_dBm < -20 {
		return Wifi{ssid, channel, frequency, signal_dBm}, true
	} else {
		return Wifi{}, false
	}
}

func NewWifiDaStringa(line string) (Wifi, bool) {
	var ssid string
	var channel int
	var frequency int
	var signal_dBm int

	if line != "" {
		fmt.Sscanf(line, "%s,%d,%d,%d", &ssid, &channel, &frequency, &signal_dBm)
		return NewWifi(ssid, channel, frequency, signal_dBm)
	}
	return Wifi{}, false
}

func (wifi Wifi) String() string {
	return fmt.Sprintf("{%s,%d,%d,%d,%v}", wifi.ssid, wifi.channel, wifi.frequency, wifi.signal_dBm, ConvertiDBinWatt(wifi.signal_dBm))
}

func ConvertiDBinWatt(signal_dBm int) float64 {
	return math.Pow(10, float64(signal_dBm)/10) / 1000
}

func PiuPotente(elenco []Wifi, banda string) int {
	var max int
	for i := 0; i < len(elenco); i++ {
		if elenco[i].frequency > 2412 && elenco[i].frequency < 2484 && banda == "2GHz" {
			if elenco[i].signal_dBm > max {
				max = elenco[i].signal_dBm
			}
		} else if elenco[i].frequency > 5035 && elenco[i].frequency < 5825 && banda == "5GHz" {
			if elenco[i].signal_dBm > max {
				max = elenco[i].signal_dBm
			}
		} else /*if (elenco[i].frequency < 5035 && elenco[i].frequency > 5825 && elenco[i].frequency < 2412 && elenco[i].frequency > 2484) && banda == ""*/ {
			if elenco[i].signal_dBm > max {
				max = elenco[i].signal_dBm
			}
		}
	}
	return max
}

func main() {
	var elenco []Wifi
	var line string
	var banda string

	if len(os.Args) > 2 {
		banda = os.Args[2]
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		if err != nil {
			break
		}
		w, _ := NewWifiDaStringa(line)
		elenco = append(elenco, w)
	}
	fmt.Println(PiuPotente(elenco, banda))
}
