### 参考 Hystrix 实现一个滑动窗口计数器。

```golang
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Rolling struct {
	sync.RWMutex

	windowSize time.Duration
	tick       time.Duration
	samples    []int64
	stopping   chan struct{}
	tail       int64
	size       int
}

func New(windowSize, tick time.Duration) (*Rolling, error) {
	if windowSize == 0 {
		return nil, errors.New("rolling window cannot be zero")
	}

	if tick == 0 {
		return nil, errors.New("tick cannot be zero")
	}

	if windowSize <= tick || windowSize%tick != 0 {
		return nil, errors.New("window size has to be a multiplier of granularity size")
	}

	r := &Rolling{
		windowSize:   windowSize,
		tick:     tick,
		samples:  make([]int64, int(windowSize/tick)),
		stopping: make(chan struct{}, 1),
	}

	go r.shifter()

	return r, nil
}

func (r *Rolling) shifter() {
	ticker := time.NewTicker(r.tick)

	for {
		select {
		case <-ticker.C:
			r.slidingRolling()
		case <-r.stopping:
			return
		}
	}
}

func (r *Rolling) slidingRolling() {
	r.Lock()
	defer r.Unlock()

	r.tail = r.tail + 1
	if r.tail >= int64(len(r.samples)) {
		r.tail = 0
	}
	r.samples[r.tail] = 0
}

func (r *Rolling) Add() {
	r.Lock()
	defer r.Unlock()
	r.samples[r.tail]++
}

func (r *Rolling) AddCount(n int64) {
	r.Lock()
	defer r.Unlock()
	r.samples[r.tail] += n
}

func (r *Rolling) Reset() {
	r.Lock()
	defer r.Unlock()

	r.tail, r.size = 0, 0
	for i := range r.samples {
		r.samples[i] = 0
	}
}

func (r *Rolling) Total() int64 {
	r.RLock()
	defer r.RUnlock()

	var total int64
	for i := range r.samples {
		total += r.samples[i]
	}
	return total
}


func main() {
	w, err := New(6*time.Second, 3*time.Second)
	if err != nil {
		fmt.Println("new rolling window error", err)
	}
	go func() {
		for {
			time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)
			//time.Sleep(5 * time.Second)
			w.AddCount(rand.Int63n(5)+3)
		}
	}()
	for i := 0; i < 20; i++ {
		fmt.Printf("%v ", w.Total())
		time.Sleep(2*time.Second)
		if i == 5 {
			w.Reset()
		}
	}
}
```
