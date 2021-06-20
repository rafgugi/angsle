package battery

var (
	alertPercentWhenCharging    = []int{90, 95, 100}
	alertPercentWhenNotCharging = []int{20, 10, 5, 0}
)

// Battery is a battery of airtel modem
type Battery struct {
	Percentage int
	IsCharging bool

	lastAlertPercentage int
}

// New creates new battery instance
func New(percentage int, isCharging bool) Battery {
	return Battery{
		Percentage:          percentage,
		IsCharging:          isCharging,
		lastAlertPercentage: percentage,
	}
}

// Update updates the current percentage and charging state
func (b *Battery) Update(percentage int, _isCharging bool) {
	if b.Percentage < percentage {
		b.IsCharging = true
	} else if b.Percentage > percentage {
		b.IsCharging = false
	}

	b.Percentage = percentage
}

// ShouldAlert will return true when it should
func (b *Battery) ShouldAlert() bool {
	// fmt.Printf("[%d%%] %v, last: %d%%\n", b.Percentage, b.IsCharging, b.lastAlertPercentage)

	if b.IsCharging {
		for _, threshold := range alertPercentWhenCharging {
			if b.Percentage < threshold || b.lastAlertPercentage >= threshold {
				continue
			}

			b.lastAlertPercentage = b.Percentage
			return true
		}
	} else {
		for _, threshold := range alertPercentWhenNotCharging {
			if b.Percentage > threshold || b.lastAlertPercentage <= threshold {
				continue
			}

			b.lastAlertPercentage = b.Percentage
			return true
		}
	}

	return false
}
