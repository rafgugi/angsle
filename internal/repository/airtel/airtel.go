package airtel

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rafgugi/angsle/battery"
)

type Airtel struct {
	host     string
	password string
	battery  *battery.Battery
}

type batteryResponse struct {
	IsCharging string `json:"charging"`
	Percentage string `json:"capacity"`
}

func New(host string, password string) *Airtel {
	var b *battery.Battery
	return &Airtel{
		host:     host,
		password: password,
		battery:  b,
	}
}

func (a Airtel) login() error {
	u, _ := url.Parse(fmt.Sprintf("%s/reqproc/proc_post?goformId=LOGIN", a.host))
	password := base64.StdEncoding.EncodeToString([]byte(a.password))

	q := u.Query()
	q.Set("password", password)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	if string(data) == "" {
		return fmt.Errorf("Response from login is empty")
	}

	return nil
}

func (a Airtel) batteryInfo() (int, bool, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/reqproc/proc_get?cmd=get_battery_info", a.host))
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, false, err
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	raw := batteryResponse{}
	if err := json.Unmarshal(data, &raw); err != nil {
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

func (a *Airtel) UpdateBattery() error {
	fmt.Println("-------------- Login --------------")
	err := a.login()
	if err != nil {
		return err
	}

	fmt.Println("-------------- Get Battery Info --------------")
	percentage, isCharging, err := a.batteryInfo()
	if err != nil {
		return err
	}
	a.battery.Update(percentage, isCharging)

	return nil
}

func (a *Airtel) GetBattery() *battery.Battery {
	fmt.Printf("battery: %+v\n", a.battery)
	return a.battery
}
