package battery

var (
	alertPercentWhenCharging    = []int{90, 95, 100}
	alertPercentWhenNotCharging = []int{20, 10, 5, 0}
)

// Battery is a battery of airtel modem
type Battery struct {
	Percentage int
	IsCharging bool
}

// Update updates the current percentage and charging state
func (b *Battery) Update(percentage int, _ bool) {
	if b.Percentage < percentage {
		b.IsCharging = true
	} else if b.Percentage > percentage {
		b.IsCharging = false
	}

	b.Percentage = percentage
}

// ShouldAlert will return true when it should
func (b Battery) ShouldAlert() bool {
	return false
}
