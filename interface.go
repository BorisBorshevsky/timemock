//go:generate mockgen -destination=mocks/clock_mock.go "github.com/BorisBorshevsky/timemock" Clock

package timemock

import (
	"sync"
	"time"
)

type Clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
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
