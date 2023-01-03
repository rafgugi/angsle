package main

import (
	"fmt"
	"time"

	"github.com/rafgugi/angsle/battery"
	"github.com/rafgugi/angsle/entity"
	"github.com/rafgugi/angsle/internal/repository/huawei"
)

const sleepTime = 30

func main() {
	var modem entity.Modem
	var b *battery.Battery

	modem = huawei.Huawei{
		Host: "http://192.168.8.1",
	}

	for {
		percentage, isCharging, err := modem.GetBattery()
		if err != nil {
			fmt.Println("error: " + err.Error())
		}

		if b == nil {
			b = battery.New(percentage, isCharging)
		}
		b.Update(percentage, isCharging)
		fmt.Printf("battery: %+v\n", *b)

		if b.ShouldAlert() {
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
