package counter

import (
	"fmt"
	"sync/atomic"
)

type Counter struct {
	value uint64
	msg   string
}

// creating a counter, and calling fn after every increase of the counter
func New(msg string) Counter {
	return Counter{
		msg: msg,
	}
}

func (c *Counter) Value() uint64 {
	return atomic.LoadUint64(&c.value)
}

func (c *Counter) Increase() {
	atomic.AddUint64(&c.value, 1)
	c.logValue()
}

func (c *Counter) logValue() {
	fmt.Printf("\r%s: %d", c.msg, c.Value())
}
