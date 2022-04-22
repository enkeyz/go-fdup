package counter

import "sync/atomic"

type Counter struct {
	value   uint64
	logFunc func(value uint64)
}

// creating a counter, and calling fn after every increase of the counter
func NewCounter(fn func(value uint64)) Counter {
	return Counter{
		logFunc: fn,
	}
}

func (c *Counter) Value() uint64 {
	return atomic.LoadUint64(&c.value)
}

func (c *Counter) Increase() {
	atomic.AddUint64(&c.value, 1)
	c.logFunc(c.value)
}
