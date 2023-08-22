package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal:=make(chan bool)

	evilNinja := "laude apna kaam kar"
	go task(evilNinja,smokeSignal)
	
	//time.Sleep(time.Second ) //because go key will create different routine
	//and it is taking 1 second of time main will be executed even before
	//task() is executed so that's why we are saying to main to wait for 3 seconds
	//before getting completed

	fmt.Println(<-smokeSignal)

}

func task(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Target is", target)
	attacked <- true
}
