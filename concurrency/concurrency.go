package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	evilNinjas := []string{"Tommy", "joohny", "bobby", "andy"}
	for _, v := range evilNinjas {
		go attack(v)
	}
	time.Sleep(time.Second * 2)//without this what will happen here is
	//go attack(v) is creating different routines 
	//for different attack which means main() is different routine as well and
	//sometimes what happens is main()func will exit even before another go routines
	//are completely executed  that is why use time.sleep for synchronize
}

func attack(target string) {
	fmt.Println("Throwing ninja starts at ", target)
	time.Sleep(time.Second)
}
