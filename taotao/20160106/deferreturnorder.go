package main

import (
	"fmt"
	"time"
)

func deferReturnOrder(i int) (x int, timer string) {

	defer func() {
		time.Sleep(time.Second * 5)
		i++
		fmt.Printf("defer2 value: %d\n defer2 time is: %s\n", i, time.Now().String())
	}()

	defer func() {
		limiter := time.Tick(time.Millisecond * 1000)
		<-limiter
		i++
		fmt.Printf("defer value: %d\n defer time is: %s\n", i, time.Now().String())
	}()

	defer fmt.Printf("de0 time %s\n", time.Now().String())

	return i, time.Now().String()
}

func main() {
	i, j := deferReturnOrder(1)
	fmt.Printf("return value: %d, %s\n", i, j)
}
