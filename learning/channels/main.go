package main

import (
	"fmt"
	"time"
)

func main() {

	// timer1 := time.NewTimer(2 * time.Second)

	<-time.After(2 * time.Second)
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(4 * time.Second)
	go func() {
		fmt.Println("before Timer 2")
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(time.Second)

}
