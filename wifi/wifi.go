package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Wifi struct {
	ssid                           string
	channel, frequency, signal_dBm int
}

func NewWifi(ssid string, channel int, frequency int, signal_dBm int) (Wifi, bool) {
	if len(ssid) > 0 {
		if (channel > 0 && channel < 15) && (frequency >= 2412 && frequency <= 2484) {
			if signal_dBm < -20 {
				return Wifi{ssid, channel, frequency, signal_dBm}, true
			}
		}
		if (channel > 6 && channel < 197) && (frequency >= 5035 && frequency <= 5980) {
			if signal_dBm < -20 {
				return Wifi{ssid, channel, frequency, signal_dBm}, true
			}
		}
	}
	return Wifi{}, false
}

func NewWifiDaStringa(line string) (Wifi, bool) {
	var ssid string
	var channel, frequency, signal_dBm int
	_, err := fmt.Sscanf(line, "%s,%d,%d,%d", &ssid, &channel, &frequency, &signal_dBm)
	if err == nil {
		return NewWifi(ssid, channel, frequency, signal_dBm)
	}
	return Wifi{}, false
}

func (wifi Wifi) String() string {
	return fmt.Sprintf("%s,%d,%d,%d", wifi.ssid, wifi.channel, wifi.frequency, wifi.signal_dBm)
}

func ConvertiDBinWatt(signal_dBm int) float64 {
	return math.Pow(10, float64(signal_dBm)/10) / 1000
}

func PiuPotente(elenco []Wifi, banda string) int {
	var max int
	for i := 0; i < len(elenco)-1; i++ {
		if banda == "2GHz" {
			if elenco[i].frequency >= 2412 && elenco[i].frequency <= 2484 {
				if elenco[i].signal_dBm > max {
					max = elenco[i].signal_dBm
				}
			}
		} else if banda == "5GHz" {
			if elenco[i].frequency >= 5035 && elenco[i].frequency <= 5980 {
				if elenco[i].signal_dBm > max {
					max = elenco[i].signal_dBm
				}
			}
		} else {
			if elenco[i].signal_dBm > max {
				max = elenco[i].signal_dBm
			}
		}
	}
	return max
}

func main() {

	var wifis []Wifi
	var banda string
	args := os.Args[1:]

	f, _ := os.Open(args[0])
	defer f.Close()

	if len(args) > 1 {
		banda = args[1]
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		w, _ := NewWifiDaStringa(line)
		wifis = append(wifis, w)
	}

	strongerIndex := PiuPotente(wifis, banda)
	strongerWifi := wifis[strongerIndex].String()
	fmt.Println(strongerWifi)
}
