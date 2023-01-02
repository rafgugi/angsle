package entity

// Modem is a device that has battery
type Modem interface {
	GetBattery() (percentage int, isCharging bool, err error)
}
