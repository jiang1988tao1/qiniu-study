package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var a int64 = 100000
	var mutex int64 = 0
	var lock int64 = 0
	var Unlock int64 = 1

	for n := 0; n < 10; n++ {

		go func() {

			for {
				if !atomic.CompareAndSwapInt64(&mutex, lock, 1) {
					if atomic.LoadInt64(&a) > 0 {
						a = atomic.AddInt64(&a, -1)

					} else {
						atomic.CompareAndSwapInt64(&mutex, Unlock, 0)
						break
					}
					atomic.CompareAndSwapInt64(&mutex, Unlock, 0)
				}
			}

		}()

	}

	time.Sleep(time.Second)
	fmt.Println(a)
}
