package main

import "sync"

type Counter struct {
	mu    sync.Mutex // mutex allows for exclusion lock, ensuring one go routine counter at a time.
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
