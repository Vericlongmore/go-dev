package main

import (
	_ "fmt"
	"sync/atomic"
	_ "time"
)

type slidingCounter struct {
	//以时间为key存储的窗口
	windows map[int64]*uint32
	//时间间隔
	interval int64
}

func main() {
	//模拟计数
	counter := NewSlidingCounter(10)
	for _, request := range []float64{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		counter.Increment(request)

	}

}

func (sc *slidingCounter) Increment(i float64) {
	if i == 0 {
		return
	}

	atomic.AddUint32(&w.Value, 1)

}
