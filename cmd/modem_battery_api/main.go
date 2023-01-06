package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rafgugi/angsle/battery"
	"github.com/rafgugi/angsle/internal/repository/huawei"
	"github.com/rafgugi/angsle/modem"
)

const (
	port      = 8411
	sleepTime = 30
)

type batteryResponse struct {
	Percentage  int  `json:"percentage"`
	IsCharging  bool `json:"is_charging"`
	ShouldAlert bool `json:"should_alert"`
}

func newBatteryResponse(b *battery.Battery) *batteryResponse {
	if b == nil {
		return nil
	}

	return &batteryResponse{
		Percentage:  b.Percentage,
		IsCharging:  b.IsCharging,
		ShouldAlert: b.ShouldAlert(),
	}
}

func main() {
	var m modem.Modem
	m = huawei.New("http://192.168.8.1")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "battery.html")
	})

	http.HandleFunc("/api/battery", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := m.UpdateBattery(); err != nil {
			fmt.Println("error: " + err.Error())
			respondError(w, err)
			return
		}

		b := newBatteryResponse(m.GetBattery())
		if b == nil {
			respondError(w, fmt.Errorf("Battery not found"))
			return
		}

		json.NewEncoder(w).Encode(b)
	})

	sigChan := make(chan bool)
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	}()
	fmt.Printf("server run with port %d\n", port)
	<-sigChan
}

func respondError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
}
