package main

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.m.Lock()
		// 使用defer原因是在执行f时panic的话可以保证锁释放
		defer o.m.Unlock()
		// 必须有该判断，虽然已经加锁，但是防止在一个goroutine在执行完atomic.StoreUint32(&o.done, 1)时进入外层if内，这样f会被执行两遍
		if o.done == 0 {
			// 使用defer原因是在执行f时panic的话可以保证f只被执行一次，并且需要放在f上方
			defer atomic.StoreUint32(&o.done, 1)
			f()
		}
	}
}
