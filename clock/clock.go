package clock

import "time"

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (r RealClocker) Now() time.Time {
	return time.Now()
}

type FixedClocker struct{}

func (c FixedClocker) Now() time.Time {
	return time.Date(2025, 1, 6, 16, 32, 0, 0, time.UTC)
}
