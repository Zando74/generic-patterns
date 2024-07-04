package main

import (
	"sync"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Decrement() {
	c.value--
}

func (c *Counter) Value() int {
	return c.value
}

type CounterProxy struct {
	counter Counter
	mu      sync.Mutex
}

func (c *CounterProxy) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter.Increment()
}

func (c *CounterProxy) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter.Decrement()
}

func MainProxyExample() {
	Counter := &Counter{}
	ProxifiedCounter := CounterProxy{counter: *Counter}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			ProxifiedCounter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	println((ProxifiedCounter.counter.Value())) // 100

}
