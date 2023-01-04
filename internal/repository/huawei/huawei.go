package huawei

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rafgugi/angsle/battery"
)

type Huawei struct {
	host    string
	battery *battery.Battery
}

type statusResponse struct {
	IsCharging string `xml:"usbup"`
	Percentage string `xml:"BatteryPercent"`
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
		return err
	}

	percentage, err := strconv.Atoi(raw.Percentage)
	if err != nil {
		return err
	}

	isCharging := false
	if raw.IsCharging == "1" {
		isCharging = true
	}
	if h.battery == nil {
		h.battery = battery.New(percentage, isCharging)
	}

	h.battery.Update(percentage, isCharging)
	return nil
}

func (h *Huawei) GetBattery() *battery.Battery {
	fmt.Printf("battery: %+v\n", h.battery)
	return h.battery
}
