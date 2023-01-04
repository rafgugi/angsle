package modem

import "github.com/rafgugi/angsle/battery"

// Modem is a device that has battery
type Modem interface {
	// UpdateBattery updates the current battery level usually by communicating
	// to the modem itself
	UpdateBattery() error

	// GetBattery gets the latest status of the battery, or return nil if the
	// battery is not initialized. It doesn't do the UpdateBattery
	GetBattery() *battery.Battery
}
