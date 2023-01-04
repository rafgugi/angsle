package battery

var (
	alertPercentWhenCharging    = 95
	alertPercentWhenNotCharging = 20
)

// Battery is a battery that has capacity up to 100%. It will be decreased
// overtime, but can be increased by charging
type Battery struct {
	Percentage int
	IsCharging bool
}

// New creates new battery instance
func New(percentage int, isCharging bool) *Battery {
	return &Battery{
		Percentage: percentage,
		IsCharging: isCharging,
	}
}

// Update updates the current percentage and charging state
func (b *Battery) Update(percentage int, isCharging bool) {
	if b.Percentage < percentage {
		b.IsCharging = true
	} else if b.Percentage > percentage {
		b.IsCharging = false
	} else {
		b.IsCharging = isCharging
	}

	b.Percentage = percentage
}

// ShouldAlert will return true when it should
func (b *Battery) ShouldAlert() bool {
	if b.IsCharging {
		if b.Percentage >= alertPercentWhenCharging {
			return true
		}
	} else {
		if b.Percentage <= alertPercentWhenNotCharging {
			return true
		}
	}

	return false
}
