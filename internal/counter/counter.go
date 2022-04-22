package counter

import (
	"sync"
)

type Counter struct {
	mu      sync.Mutex
	value   int64
	logFunc func(value int64)
}

// creating a counter, and calling fn after after increase of the counter
func NewCounter(fn func(value int64)) Counter {
	return Counter{
		logFunc: fn,
	}
}

func (c *Counter) Value() int64 {
	return c.value
}

func (c *Counter) Increase() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++

	c.logFunc(c.value)
}
