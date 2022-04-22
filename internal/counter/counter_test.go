package counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times", func(t *testing.T) {
		counter := NewCounter(func(counterValue uint64) {})
		counter.Increase()
		counter.Increase()
		counter.Increase()

		if counter.Value() != 3 {
			t.Errorf("got %d, but expected %d", counter.Value(), 3)
		}
	})
}
