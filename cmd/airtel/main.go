package main

import (
	"fmt"
	"time"

	"github.com/rafgugi/angsle/battery"
)

func main() {
	a := Airtel{
		host:     "http://192.168.8.1",
		password: "admin",
	}

	var b *battery.Battery

	currentCharging := false
	shouldAlert := false
	for {
		fmt.Println("-------------- Login --------------")
		err := a.Login()
		if err != nil {
			fmt.Println("error: " + err.Error())
		}

		fmt.Println("-------------- Get Battery Info --------------")
		percentage, isCharging, err := a.BatteryInfo()
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
		time.Sleep(30 * time.Second)
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
