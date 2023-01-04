package battery_test

import (
	"testing"

	"github.com/rafgugi/angsle/battery"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	type batteryCase struct {
		percentage int
		isCharging bool
		want       bool
	}
	tests := map[string]struct {
		initial  batteryCase
		overtime []batteryCase
	}{
		"discharging": {
			initial: batteryCase{100, false, false},
			overtime: []batteryCase{
				{100, false, false},
				{99, false, false},
				{21, false, false},
				{21, false, false},
				{20, false, false},
				{19, false, false},
				{19, false, false},
				{5, false, false},
				{0, false, false},
				{0, false, false},
			},
		},
		"charging": {
			initial: batteryCase{0, false, false},
			overtime: []batteryCase{
				{0, true, true},
				{0, true, true},
				{5, true, true},
				{9, true, true},
				{95, true, true},
				{95, true, true},
				{96, true, true},
				{100, true, true},
			},
		},
		"discharging but increasing": {
			initial: batteryCase{100, false, false},
			overtime: []batteryCase{
				{100, false, false},
				{99, false, false},
				{21, false, false},
				{21, false, false},
				{20, false, false},
				{19, false, false},
				{19, false, false},
				{20, false, true},
				{20, false, false},
				{5, false, false},
				{0, false, false},
				{0, false, false},
			},
		},
		"charging but decreasing": {
			initial: batteryCase{0, false, false},
			overtime: []batteryCase{
				{0, true, true},
				{0, true, true},
				{5, true, true},
				{9, true, true},
				{95, true, true},
				{94, true, false},
				{94, true, true},
				{95, true, true},
				{96, true, true},
				{100, true, true},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			b := battery.New(tc.initial.percentage, tc.initial.isCharging)

			for _, tt := range tc.overtime {
				b.Update(tt.percentage, tt.isCharging)
				got := b.IsCharging
				assert.Equal(t, tt.percentage, b.Percentage)
				assert.Equal(t, tt.want, got, "[%d%%] when %s should %s",
					tt.percentage,
					map[bool]string{true: "charging", false: "not charging"}[tt.isCharging],
					map[bool]string{true: "charging", false: "not charging"}[tt.want],
				)
			}
		})
	}
}

func TestShouldAlert(t *testing.T) {
	type batteryCase struct {
		percentage  int
		isCharging  bool
		shouldAlert bool
	}
	tests := map[string]struct {
		initial  batteryCase
		overtime []batteryCase
	}{
		"discharging": {
			initial: batteryCase{100, false, false},
			overtime: []batteryCase{
				{100, false, false},
				{99, false, false},
				{21, false, false},
				{20, false, true},
				{19, false, true},
				{19, false, true},
				{19, false, true},
				{16, false, true},
				{15, false, true},
				{11, false, true},
				{11, false, true},
				{9, false, true},
				{9, false, true},
				{5, false, true},
				{0, false, true},
			},
		},
		"charging": {
			initial: batteryCase{0, false, false},
			overtime: []batteryCase{
				{0, true, false},
				{5, true, false},
				{9, true, false},
				{19, true, false},
				{20, true, false},
				{21, true, false},
				{89, true, false},
				{90, true, false},
				{91, true, false},
				{92, true, false},
				{95, true, true},
				{95, true, true},
				{96, true, true},
				{100, true, true},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			b := battery.New(tc.initial.percentage, tc.initial.isCharging)

			for _, tt := range tc.overtime {
				b.Update(tt.percentage, tt.isCharging)
				ok := b.ShouldAlert()
				assert.Equal(t, tt.shouldAlert, ok, "[%d%%] when %s should %s",
					tt.percentage,
					map[bool]string{true: "charging", false: "not charging"}[tt.isCharging],
					map[bool]string{true: "alert", false: "not alert"}[tt.shouldAlert],
				)
			}
		})
	}
}
