package timemock

import (
	"sync"
	"time"
)

var (
	now = time.Now
)

type timemockClock struct {
	rw         *sync.RWMutex
	frozen     bool
	traveled   bool
	freezeTime time.Time
	travelTime time.Time
	scale      float64
}

func (c *timemockClock) Scale(scale float64) {
	c.scale = scale
}

func (c *timemockClock) Now() time.Time {
	if c.frozen {
		c.rw.RLock()
		defer c.rw.RUnlock()
		return c.freezeTime
	}

	if c.traveled {
		c.rw.RLock()
		defer c.rw.RUnlock()
		return c.freezeTime.Add(time.Duration(float64(time.Since(c.travelTime)) * c.scale))
	}

	return now()
}

func (c *timemockClock) Freeze(t time.Time) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.freezeTime = t
	c.frozen = true
}

func (c *timemockClock) Travel(t time.Time) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.freezeTime = t
	c.travelTime = time.Now()
	c.traveled = true
}

func (c *timemockClock) Since(t time.Time) time.Duration {
	return c.Now().Sub(t)
}

func (c *timemockClock) Return() {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.frozen = false
	c.traveled = false
	c.scale = 1
}
