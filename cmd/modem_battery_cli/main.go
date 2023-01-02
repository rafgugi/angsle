package main

import (
	"fmt"
	"time"

	"github.com/rafgugi/angsle/battery"
	"github.com/rafgugi/angsle/entity"
	"github.com/rafgugi/angsle/internal/repository/airtel"
)

const sleepTime = 30

func main() {
	var modem entity.Modem
	var b *battery.Battery

	modem = airtel.Airtel{
		Host:     "http://192.168.8.1",
		Password: "admin",
	}

	currentCharging := false
	shouldAlert := false
	for {
		percentage, isCharging, err := modem.GetBattery()
		if err != nil {
			fmt.Println("error: " + err.Error())
		}

		if b == nil {
			b = battery.New(percentage, isCharging)
		}

		// reset alert if charge state updated
		if currentCharging != isCharging {
			shouldAlert = false
			currentCharging = isCharging
		}

		b.Update(percentage, isCharging)
		fmt.Printf("battery: %+v\n", *b)

		if b.ShouldAlert() || shouldAlert {
			shouldAlert = true
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
