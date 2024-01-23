package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Wifi struct {
	ssid                           string
	channel, frequency, signal_dBm int
}

func NewWifi(ssid string, channel int, frequency int, signal_dBm int) (Wifi, bool) {
	if len(ssid) > 0 && signal_dBm <= -20 {
		if (channel >= 1 && channel <= 14) && (frequency >= 2412 && frequency <= 2484) {
			return Wifi{ssid, channel, frequency, signal_dBm}, true
		} else if (channel >= 7 && channel <= 196) && (frequency >= 5035 && frequency <= 5980) {
			return Wifi{ssid, channel, frequency, signal_dBm}, true
		}
	}
	return Wifi{}, false
}

func NewWifiDaStringa(line string) (Wifi, bool) {
	fields := strings.Split(line, ",")
	ssid := fields[0]
	channel, _ := strconv.Atoi(fields[1])
	frequency, _ := strconv.Atoi(fields[2])
	signal_dBm, _ := strconv.Atoi(fields[3])
	return NewWifi(ssid, channel, frequency, signal_dBm)
}

func (wifi Wifi) String() string {
	return fmt.Sprintf("%s,%d,%d,%d,%f", wifi.ssid, wifi.channel, wifi.frequency, wifi.signal_dBm, ConvertiDBinWatt(wifi.signal_dBm))
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

	var counter int
	var elenco []Wifi
	var banda string

	f, _ := os.Open(os.Args[1])
	defer f.Close()

	if len(os.Args) > 2 {
		banda = os.Args[2]
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		wifi, _ := NewWifiDaStringa(line)

		if counter > 0 {
			elenco = append(elenco, wifi)
		}
		counter++
	}

	fmt.Println(elenco[PiuPotente(elenco, banda)].String())
}
