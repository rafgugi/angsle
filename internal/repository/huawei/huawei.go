package huawei

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Huawei struct {
	Host string
}

type statusResponse struct {
	IsCharging string `xml:"usbup"`
	Percentage string `xml:"BatteryPercent"`
}

func (h Huawei) getStatus() (statusResponse, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/api/monitoring/status", h.Host))
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

func (h Huawei) GetBattery() (int, bool, error) {
	fmt.Println("-------------- Get Status --------------")
	raw, err := h.getStatus()
	if err != nil {
		return 0, false, err
	}

	percentage, err := strconv.Atoi(raw.Percentage)
	if err != nil {
		return 0, false, err
	}

	isCharging := false
	if raw.IsCharging == "1" {
		isCharging = true
	}

	return percentage, isCharging, nil
}
