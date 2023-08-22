package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)

	go throwingNinjaStar(channel)
	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}

}

func throwingNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	num := 3
	for i := 0; i < num; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprintf("your scored: ", score)
	}
	close(channel)
}
/*
1.when throwingNinjaStar() for loop inside this execute it will fill up channel
with one output and after that it cannot because it is blocking operation
but as soon as main go routine fetches info from channel the throwingNinjaStar()
will now become non blocking go routine which will can now insert new value

2.as soon as main for loop fetches the value from channel it will try to fetch
another value but it cannot because channel is empty and it needs to be 
updated by throwingNinjaStar() go routine but after throwingNinjaStar()
filling up channel it can fetch data from channel

3. the above every point was in contrast with normal for loop but when we use
range the output changes majorly bcoz for message := channel the message
will get all the value from channel and it will expect more values from channel
like normal for loop only but as message has extracted everything from channel in
one go it cannot that's why we use "close(channel)" which will help for range
that this channel is no longer in use as soon as all the messages are extracted
out u should also exit out of for loop when u are using range keyword with for.
*/