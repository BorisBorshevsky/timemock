package timemock

import (
	"sync"
	"time"
)

//go:generate mockgen -destination=internal/clock_mock.go . Clock
type Clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
	Until(time.Time) time.Duration
	Freeze(time.Time)
	Travel(time.Time)
	Scale(float64)
	Return()
}

func New() Clock {
	return &timemockClock{
		scale: 1,
		rw:    new(sync.RWMutex),
	}
}
