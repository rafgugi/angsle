package main

import (
	"fmt"
	"time"

	"github.com/rafgugi/angsle/internal/repository/huawei"
	"github.com/rafgugi/angsle/modem"
)

const sleepTime = 30

func main() {
	var m modem.Modem
	m = huawei.New("http://192.168.8.1")

	for {
		if err := m.UpdateBattery(); err != nil {
			fmt.Println("error: " + err.Error())
		}

		b := m.GetBattery()
		if b != nil && b.ShouldAlert() {
			alert()
		}

		fmt.Println("-------------- Sleep --------------")
		time.Sleep(sleepTime * time.Second)
	}
}

func alert() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(" ▄▄▄       ██▓    ▓█████  ██▀███  ▄▄▄█████▓")
	fmt.Println("▒████▄    ▓██▒    ▓█   ▀ ▓██ ▒ ██▒▓  ██▒ ▓▒")
	fmt.Println("▒██  ▀█▄  ▒██░    ▒███   ▓██ ░▄█ ▒▒ ▓██░ ▒░")
	fmt.Println("░██▄▄▄▄██ ▒██░    ▒▓█  ▄ ▒██▀▀█▄  ░ ▓██▓ ░ ")
	fmt.Println(" ▓█   ▓██▒░██████▒░▒████▒░██▓ ▒██▒  ▒██▒ ░ ")
	fmt.Println(" ▒▒   ▓▒█░░ ▒░▓  ░░░ ▒░ ░░ ▒▓ ░▒▓░  ▒ ░░   ")
	fmt.Println("  ▒   ▒▒ ░░ ░ ▒  ░ ░ ░  ░  ░▒ ░ ▒░    ░    ")
	fmt.Println("  ░   ▒     ░ ░      ░     ░░   ░   ░      ")
	fmt.Println("      ░  ░    ░  ░   ░  ░   ░              ")
	fmt.Println("")
	fmt.Println("")
}
