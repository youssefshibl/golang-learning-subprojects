package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

// time.After(// 2 * time.Second) vs // time.NewTimer(2 * time.Second)
// time.After is a function that returns a channel that will send the current time after a specified duration.
// time.NewTimer is a struct that represents a timer that will send the current time after a specified duration.

// time.NewTimer(// 2 * time.Second) vs // time.sleep(2 * time.Second)
// time.NewTimer can stop the timer before it fires, while time.Sleep cannot.
// time.Sleep can't be stopped, it blocks the current goroutine until the duration has passed.
