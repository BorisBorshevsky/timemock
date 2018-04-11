package timemock

import (
	"sync"
	"time"
)

type Clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
	Freeze(time.Time)
	Scale(float64)
}

func New() Clock {
	return &timemockClock{
		scale: 1,
		rw:    new(sync.RWMutex),
	}
}
