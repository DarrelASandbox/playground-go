package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

/*
  Any goroutine calling Inc will acquire the lock on Counter if they are first.
  All the other goroutines will have to wait for it to be Unlocked before getting access.
*/

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
