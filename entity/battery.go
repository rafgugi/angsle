package entity

// Battery is a battery that has capacity up to 100%. It will be decreased
// overtime, but can be increased by charging
type Battery interface {
	Update(percentage int, isCharging bool)
	ShouldAlert() bool
}
