package huawei

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/rafgugi/angsle/battery"
	"github.com/rafgugi/angsle/modem"
)

type Huawei struct {
	host    string
	battery *battery.Battery
}

type statusResponse struct {
	IsCharging string `xml:"usbup"`
	Percentage int    `xml:"BatteryPercent"`
}

func isModem() {
	var _ modem.Modem = (*Huawei)(nil)
}

func New(host string) *Huawei {
	var b *battery.Battery
	return &Huawei{
		host:    host,
		battery: b,
	}
}

func (h Huawei) getStatus() (statusResponse, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/api/monitoring/status", h.host))
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("Update-Cookie", "UpdateCookie")

	raw := statusResponse{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return raw, err
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	if err := xml.Unmarshal(data, &raw); err != nil {
		return raw, err
	}

	return raw, nil
}

func (h *Huawei) UpdateBattery() error {
	fmt.Println("-------------- Get Status --------------")
	raw, err := h.getStatus()
	if err != nil {
		h.battery = nil
		return err
	}

	isCharging := false
	if raw.IsCharging == "1" {
		isCharging = true
	}
	if h.battery == nil {
		h.battery = battery.New(raw.Percentage, isCharging)
	} else {
		h.battery.Update(raw.Percentage, isCharging)
	}

	return nil
}

func (h *Huawei) GetBattery() *battery.Battery {
	fmt.Printf("battery: %+v\n", h.battery)
	return h.battery
}
