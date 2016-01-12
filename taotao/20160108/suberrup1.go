package main

import (
	"fmt"
	"sync"
	// "time"
)

var m *sync.Mutex

func main() {
	var a int64 = 100000
	m = new(sync.Mutex)
	var wg sync.WaitGroup
	wg.Add(10)
	for n := 0; n < 10; n++ {

		go func() {

			for {
				m.Lock()
				if a > 0 {
					a--
				} else {
					m.Unlock()
					break
				}
				m.Unlock()
			}

			wg.Done()
		}()

	}

	wg.Wait()
	// time.Sleep(time.Second)
	fmt.Println(a)

}
