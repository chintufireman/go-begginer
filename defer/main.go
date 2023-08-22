package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")
	mydefer()
}

//after using defer keyword ur statement goes to the the end of the file
//and every statement having defer is executed LIFO order that's why two, one , world
func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
