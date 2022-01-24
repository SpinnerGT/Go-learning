package main

import "fmt"

func main1() {
	var myInt int
	var myIntPointer *int

	myInt = 42
	myIntPointer = &myInt

	fmt.Println("the value of myIntPointer is", *myIntPointer)

}

func main() {

}
