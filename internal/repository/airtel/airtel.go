package airtel

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Airtel struct {
	Host     string
	Password string
}

type batteryResponse struct {
	IsCharging string `json:"charging"`
	Percentage string `json:"capacity"`
}

func (a Airtel) login() error {
	u, _ := url.Parse(fmt.Sprintf("%s/reqproc/proc_post?goformId=LOGIN", a.Host))
	password := base64.StdEncoding.EncodeToString([]byte(a.Password))

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
	u, _ := url.Parse(fmt.Sprintf("%s/reqproc/proc_get?cmd=get_battery_info", a.Host))
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

func (a Airtel) GetBattery() (int, bool, error) {
	fmt.Println("-------------- Login --------------")
	err := a.login()
	if err != nil {
		return 0, false, err
	}

	fmt.Println("-------------- Get Battery Info --------------")
	percentage, isCharging, err := a.batteryInfo()
	if err != nil {
		return 0, false, err
	}

	return percentage, isCharging, nil
}