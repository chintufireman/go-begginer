package main

import "fmt"

func main()  {
	channel:=make(chan string,2)
	//channel <- "Attacked" this is deadlock condition because 
	//the size of channel is 0 yes, and when u try to fetch from chan value from 
	//same routine it cannot be retrieved because there is still statement remainig
	//to be executed and hence thats why we use buffered channel where we
	//increase the size of channel from 0 to 1 
	//1. whene we create a channel make sure there is another routine which
	//will accept the value from other routine because channel capacity is 
	//always zero.
	//2.if u want to save it in any channel then simply create a buffered chan
	//where even if u dont have any other go routine u can simply keep it in
	//buffer


	channel <- "Attacked"
	channel <- "Not Attacked" //increase size to two from 1
	fmt.Println(<-channel)
	//output will be "Attacked" because channels are FIFO
}