package pkg

import "sync/atomic"

type safeCounter int32

func (c *safeCounter) Increment() {
	atomic.AddInt32((*int32)(c), 1)
}

func (c *safeCounter) Decrement() {
	atomic.AddInt32((*int32)(c), -1)
}

func (c *safeCounter) Set(value int32) {
	atomic.StoreInt32((*int32)(c), value)
}

func (c *safeCounter) GetValue() int32 {
	value := atomic.LoadInt32((*int32)(c))
	return value
}
